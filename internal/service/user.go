package service

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/dao"
	"github.com/go-nunu/nunu-layout-advanced/internal/model"
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/convert"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   string `json:"avatar"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type UserService struct {
	userDao *dao.UserDao
	*Service
}

func NewUserService(service *Service, userDao *dao.UserDao) *UserService {
	return &UserService{
		userDao: userDao,
		Service: service,
	}
}

func (s *UserService) Register(req *RegisterRequest) error {
	// 生成用户ID
	userId, err := s.generateUserId()
	if err != nil {
		return errors.Wrap(err, "failed to generate user ID")
	}

	// 检查用户名是否已存在
	if user, err := s.userDao.GetUserByUsername(req.Username); err == nil && user != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	// 创建用户
	user := &model.User{
		UserId:   userId,
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	if err = s.userDao.CreateUser(user); err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (s *UserService) Login(req *LoginRequest) (string, error) {
	user, err := s.userDao.GetUserByUsername(req.Username)
	if err != nil || user == nil {
		return "", errors.Wrap(err, "failed to get user by username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.Wrap(err, "failed to hash password")
	}
	// 生成JWT token
	token, err := s.generateToken(user.UserId)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}

	return token, nil
}

func (s *UserService) GetProfile(userId string) (*model.User, error) {
	user, err := s.userDao.GetUserById(userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return user, nil
}

func (s *UserService) UpdateProfile(userId string, req *UpdateProfileRequest) error {
	user, err := s.userDao.GetUserById(userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user by ID")
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userDao.UpdateUser(user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (s *UserService) generateUserId() (string, error) {
	// 生成分布式ID
	id, err := s.sonyflake.NextID()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate sonyflake ID")
	}

	// 将ID转换为字符串
	return convert.IntToBase62(int(id)), nil
}

func (s *UserService) generateToken(userId string) (string, error) {
	// 生成JWT token
	s.jwt.GenToken(userId, time.Now().Add(time.Hour*24*90))
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte("secret"))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}

	return token, nil
}
