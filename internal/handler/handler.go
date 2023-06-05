package handler

import (
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/sony/sonyflake"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger, sf *sonyflake.Sonyflake) *Handler {
	return &Handler{
		logger: logger,
	}
}
