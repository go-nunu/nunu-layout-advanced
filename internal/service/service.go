package service

import "github.com/go-nunu/nunu-layout-advanced/pkg/log"

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
