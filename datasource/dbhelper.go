package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"superstar/conf"
	"sync"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster error=", err)
		return nil
	}
	// 性能优化的时候才考虑，加上本机的SQL缓存
	cacher := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(false)
	engine.SetTZLocation(conf.SysTimeLocation)

	masterEngine = engine
	return masterEngine

}

func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driverSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceSlave error=", err)
		return nil
	} else {
		masterEngine = engine
		return slaveEngine
	}
}
