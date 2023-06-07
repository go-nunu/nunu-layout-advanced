package service

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/middleware"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/sony/sonyflake"
)

type Service struct {
	logger    *log.Logger
	sonyflake *sonyflake.Sonyflake
	jwt       *middleware.JWT
}

func NewService(logger *log.Logger, sonyflake *sonyflake.Sonyflake, jwt *middleware.JWT) *Service {
	return &Service{
		logger:    logger,
		sonyflake: sonyflake,
		jwt:       jwt,
	}
}
