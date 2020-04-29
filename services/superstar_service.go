package services

import (
	"superstar/dao"
	"superstar/datasource"
	"superstar/models"
)

// UserService handles CRUID operations of a user datamodel,
// it depends on a user repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type SuperstarService interface {
	GetAll() []models.StarInfo
	Search(country string) []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() SuperstarService {
	return &superstarService{
		dao: dao.NewSuperstarDao(datasource.InstanceMaster()),
	}
}

func (s *superstarService)GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarService)Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}

func (s *superstarService)Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *superstarService)Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *superstarService)Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *superstarService)Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
