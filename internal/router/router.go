package router

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/handler"
	"github.com/go-nunu/nunu-layout-advanced/pkg/jwt"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
)

type RouterDeps struct {
	Logger      *log.Logger
	JWT         *jwt.JWT
	UserHandler *handler.UserHandler
}
