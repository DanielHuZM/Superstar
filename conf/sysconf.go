package conf

import "time"

const SysTimeform string = "2006-01-02 15:04:05"
const SysTimeformShort string = "2006-01-02"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
