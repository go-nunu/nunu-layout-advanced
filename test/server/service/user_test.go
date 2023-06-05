package service

import (
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/internal/dao"
	"github.com/go-nunu/nunu-layout-advanced/internal/model"
	"github.com/go-nunu/nunu-layout-advanced/internal/service"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

var userService *service.UserService

func TestMain(m *testing.M) {
	fmt.Println("begin")

	os.Setenv("APP_CONF", "../../../config/local.yml")

	conf := config.NewConfig()

	logger := log.NewLog(conf)
	db := dao.NewDB(conf)
	rdb := dao.NewRedis(conf)

	srv := service.NewService(logger)
	repo := dao.NewDao(db, rdb, logger)
	userDao := dao.NewUserDao(repo)
	userService = service.NewUserService(srv, userDao)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}
func TestGetUserByEmail(t *testing.T) {
	_, err := userService.GetUserById(0)
	assert.Equal(t, err, gorm.ErrRecordNotFound, "they should be equal")
}

func TestCreateUser(t *testing.T) {
	_, err := userService.CreateUser(&model.User{
		Username: "test",
		Email:    "nunu@mail.com",
	})
	assert.NotEqual(t, err, nil, "they should be equal")
}
