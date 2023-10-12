package main

import (
	"context"
	"github.com/go-nunu/nunu-layout-advanced/cmd/migration/wire"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
