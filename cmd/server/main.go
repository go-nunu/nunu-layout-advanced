package main

import (
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/cmd/server/wire"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/http"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	servers, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", "http://localhost:"+conf.GetString("http.port")))

	//servers.
	http.Run(servers.ServerHTTP, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()

}
