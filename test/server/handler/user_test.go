package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/cmd/server/wire"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var headers = map[string]string{
	"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiJ5aHM2SGVzZmdGIiwiZXhwIjoxNjkzOTE0ODgwLCJuYmYiOjE2ODYxMzg4ODAsImlhdCI6MTY4NjEzODg4MH0.NnFrZFgc_333a9PXqaoongmIDksNvQoHzgM_IhJM4MQ",
}

func TestMain(m *testing.M) {
	fmt.Println("begin")

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewRequest(method, path string, header map[string]string, body io.Reader) (*Response, error) {
	// 测试时需要定义好 gin 的路由定义函数
	os.Setenv("APP_CONF", "../../../config/local.yml")
	conf := config.NewConfig()
	logger := log.NewLog(conf)
	logger.Info("start")

	app, _, err := wire.NewApp(conf, logger)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(method, path, body)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if strings.ToUpper(method) != "GET" && body != nil {
		req.Header.Set("Content-Type", "application/json")

	}
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	response := new(Response)
	err = json.Unmarshal([]byte(w.Body.String()), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func TestGetProfile(t *testing.T) {
	response, err := NewRequest("GET",
		fmt.Sprintf("/user"),
		headers,
		nil,
	)

	t.Log("response")
	assert.Nil(t, err)
	assert.Equal(t, 1, response.Code)
}
func TestUpdateProfile(t *testing.T) {
	params, err := json.Marshal(map[string]interface{}{
		"email":    "5303221@gmail.com",
		"username": "user1",
		"nickname": "8888",
	})
	assert.Nil(t, err)
	response, err := NewRequest("PUT",
		"/user",
		headers,
		bytes.NewBuffer(params),
	)

	t.Log("响应结果")
	assert.Nil(t, err)
	//assert.NotEmpty(t, response.Data)
	assert.Equal(t, 0, response.Code)
	//tsms.SendSMS2("MotokApp", "18502100065", "1234")
}
