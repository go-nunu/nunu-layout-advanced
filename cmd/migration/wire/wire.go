//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/migration"
	"github.com/go-nunu/nunu-layout-advanced/internal/provider"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// wire.go 初始化模块
func NewApp(*viper.Viper, *log.Logger) (*migration.Migrate, func(), error) {
	//log.Info("NewApp")
	panic(wire.Build(
		provider.DaoSet,
		provider.MigrateSet,
	))
}
