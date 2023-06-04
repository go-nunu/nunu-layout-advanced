package main

import (
	"github.com/go-nunu/nunu-layout/cmd/migration/wire"
	"github.com/go-nunu/nunu-layout/pkg/config"
	"github.com/go-nunu/nunu-layout/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()

}
