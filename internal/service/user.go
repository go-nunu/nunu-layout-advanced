package service

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/dao"
	"github.com/go-nunu/nunu-layout-advanced/internal/model"
)

type UserService struct {
	*Service
	userDao *dao.UserDao
}

func NewUserService(service *Service, userDao *dao.UserDao) *UserService {
	return &UserService{
		Service: service,
		userDao: userDao,
	}
}

func (s *UserService) GetUserById(id int64) (*model.User, error) {
	return s.userDao.FirstById(id)
}
func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	return s.userDao.CreateUser(user)
}
