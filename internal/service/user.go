package service

import (
	"github.com/go-nunu/nunu-layout/internal/model"
	"github.com/go-nunu/nunu-layout/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUserById(id int64) (*model.User, error) {
	return s.userRepository.FirstById(id)
}
func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	return s.userRepository.CreateUser(user)
}
