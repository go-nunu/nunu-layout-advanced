//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/job"
	"github.com/go-nunu/nunu-layout-advanced/internal/provider"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// wire.go 初始化模块
func NewApp(*viper.Viper, *log.Logger) (*job.Job, func(), error) {
	panic(wire.Build(
		provider.JobSet,
	))
}
