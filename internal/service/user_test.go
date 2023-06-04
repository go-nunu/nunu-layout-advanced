package service

import (
	"fmt"
	"github.com/go-nunu/nunu-layout/internal/repository"
	"github.com/go-nunu/nunu-layout/pkg/config"
	"github.com/go-nunu/nunu-layout/pkg/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

var userService *UserService

func TestMain(m *testing.M) {
	fmt.Println("begin")

	os.Setenv("APP_CONF", "../../config/local.yml")
	userRepository := repository.NewUserRepository(db.NewDB(config.NewConfig()))
	userService = NewUserService(userRepository)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}
func TestGetUserByEmail(t *testing.T) {
	_, err := userService.GetUserByEmail("abc")
	assert.Equal(t, err, gorm.ErrRecordNotFound, "they should be equal")
}
