package main

import (
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/cmd/server/wire"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/http"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
)

// @title           Nunu Example API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	servers, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}

	logger.Info("server start", zap.String("host", "http://localhost:"+conf.GetString("http.port")))
	logger.Info("docs addr", zap.String("addr", fmt.Sprintf("http://127.0.0.1:%d/swagger/index.html", conf.GetInt("http.port"))))

	//go grpc.Run(servers.ServerGRPC, conf)
	http.Run(servers.ServerHTTP, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()

}
