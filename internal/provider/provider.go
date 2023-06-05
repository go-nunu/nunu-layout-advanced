package provider

import (
	"github.com/go-nunu/nunu-layout/internal/dao"
	"github.com/go-nunu/nunu-layout/internal/handler"
	"github.com/go-nunu/nunu-layout/internal/job"
	"github.com/go-nunu/nunu-layout/internal/middleware"
	"github.com/go-nunu/nunu-layout/internal/migration"
	"github.com/go-nunu/nunu-layout/internal/server"
	"github.com/go-nunu/nunu-layout/internal/service"
	"github.com/go-nunu/nunu-layout/pkg/helper/sonyflake"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var SonyflakeSet = wire.NewSet(sonyflake.NewSonyflake)

var MigrateSet = wire.NewSet(migration.NewMigrate)

var JobSet = wire.NewSet(job.NewJob)

var JwtSet = wire.NewSet(middleware.NewJwt)

var DaoSet = wire.NewSet(
	dao.NewDB,
	dao.NewRedis,
	dao.NewDao,
	dao.NewUserDao,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)
