package provider

import (
	"github.com/go-nunu/nunu-layout/internal/database"
	"github.com/go-nunu/nunu-layout/internal/handler"
	"github.com/go-nunu/nunu-layout/internal/job"
	"github.com/go-nunu/nunu-layout/internal/middleware"
	"github.com/go-nunu/nunu-layout/internal/repository"
	"github.com/go-nunu/nunu-layout/internal/server"
	"github.com/go-nunu/nunu-layout/internal/service"
	"github.com/go-nunu/nunu-layout/pkg/db"
	"github.com/go-nunu/nunu-layout/pkg/sonyflake"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)
var DBSet = wire.NewSet(db.NewDB)
var SonyflakeSet = wire.NewSet(sonyflake.NewSonyflake)
var RepositorySet = wire.NewSet(repository.NewUserRepository)
var ServiceSet = wire.NewSet(service.NewUserService)
var MigrateSet = wire.NewSet(database.NewMigrate)
var JobSet = wire.NewSet(job.NewJob)
var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(
	handler.NewUserHandler,
)
