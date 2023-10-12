package service

import (
	"context"
	v1 "github.com/go-nunu/nunu-layout-advanced/api/v1"
	"github.com/go-nunu/nunu-layout-advanced/internal/model"
	"github.com/go-nunu/nunu-layout-advanced/internal/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check username
	if user, err := s.userRepo.GetByUsername(ctx, req.Username); err == nil && user != nil {
		return v1.ErrUsernameAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}
	// Generate user ID
	userId, err := s.sid.GenString()
	if err != nil {
		return errors.Wrap(err, "failed to generate user ID")
	}
	// Create a user
	user := &model.User{
		UserId:   userId,
		Username: req.Username,
		Nickname: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	if err = s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.Wrap(err, "failed to hash password")
	}
	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return &v1.GetProfileResponseData{
		UserId:   user.UserId,
		Nickname: user.Nickname,
		Username: user.Username,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user by ID")
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
