package log

import (
	"flag"
	"fmt"
	"github.com/go-nunu/nunu-layout/pkg/config"
	"go.uber.org/zap"
	"os"
	"testing"
)

var logTest *Logger

func TestMain(m *testing.M) {
	fmt.Println("begin")
	flag.Set("conf", "../../config/local.yml")
	//os.Setenv("APP_CONF", "../../config/local.yml")
	logTest = NewLog(config.NewConfig())

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}
func TestLog(t *testing.T) {
	logTest.Info("test log", zap.String("data", "test data"))
}
