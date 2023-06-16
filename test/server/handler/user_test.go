package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/internal/handler"
	"github.com/go-nunu/nunu-layout-advanced/mocks/service"

	"github.com/go-nunu/nunu-layout-advanced/internal/server"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/internal/middleware"
	"github.com/go-nunu/nunu-layout-advanced/internal/model"
	"github.com/go-nunu/nunu-layout-advanced/internal/service"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	userId = "yhs6HesfgF"

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiJ5aHM2SGVzZmdGIiwiZXhwIjoxNjkzOTE0ODgwLCJuYmYiOjE2ODYxMzg4ODAsImlhdCI6MTY4NjEzODg4MH0.NnFrZFgc_333a9PXqaoongmIDksNvQoHzgM_IhJM4MQ"
)
var hdl *handler.Handler

func TestMain(m *testing.M) {
	fmt.Println("begin")
	os.Setenv("APP_CONF", "../../../config/local.yml")
	conf := config.NewConfig()

	logger := log.NewLog(conf)
	hdl = handler.NewHandler(logger)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := service.RegisterRequest{
		Username: "xxx",
		Password: "123456",
		Email:    "xxx@gmail.com",
	}

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().Register(gomock.Any(), &params).Return(nil)

	router := setupRouter(mockUserService)
	paramsJson, _ := json.Marshal(params)

	resp := performRequest(router, "POST", "/register", bytes.NewBuffer(paramsJson))

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func TestUserHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := service.LoginRequest{
		Username: "xxx",
		Password: "123456",
	}

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().Login(gomock.Any(), &params).Return(token, nil)

	router := setupRouter(mockUserService)
	paramsJson, _ := json.Marshal(params)

	resp := performRequest(router, "POST", "/login", bytes.NewBuffer(paramsJson))

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func TestUserHandler_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().GetProfile(gomock.Any(), userId).Return(&model.User{
		Id:       1,
		UserId:   userId,
		Username: "xxxxx",
		Nickname: "xxxxx",
		Password: "xxxxx",
		Email:    "xxxxx@gmail.com",
	}, nil)

	router := setupRouter(mockUserService)
	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func TestUserHandler_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := service.UpdateProfileRequest{
		Nickname: "alan",
		Email:    "alan@gmail.com",
		Avatar:   "xxx",
	}

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().UpdateProfile(gomock.Any(), userId, &params).Return(nil)

	router := setupRouter(mockUserService)
	paramsJson, _ := json.Marshal(params)

	req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(paramsJson))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}

func setupRouter(mockUserService *mock_service.MockUserService) *gin.Engine {
	conf := config.NewConfig()
	logger := log.NewLog(conf)
	jwt := middleware.NewJwt(conf)
	userHandler := handler.NewUserHandler(hdl, mockUserService)
	gin.SetMode(gin.TestMode)
	router := server.NewServerHTTP(logger, jwt, userHandler)
	return router
}

func performRequest(r http.Handler, method, path string, body *bytes.Buffer) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}
