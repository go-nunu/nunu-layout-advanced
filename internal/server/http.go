package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout/internal/handler"
	"github.com/go-nunu/nunu-layout/internal/middleware"
	"github.com/go-nunu/nunu-layout/pkg/log"
	"github.com/go-nunu/nunu-layout/pkg/resp"
)

func NewServerHTTP(
	log *log.Logger,
	jwt *middleware.JWT,
	userHandler *handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.RequestLogMiddleware(log),
		middleware.CORSMiddleware(),
		middleware.NoAuth(log),
		middleware.ResponseLogMiddleware(log),
		//middleware.SignMiddleware(log),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})

	noAuthRouter := r.Use(middleware.NoAuth(log))
	{
		noAuthRouter.GET("/user", userHandler.GetUserById)

	}
	// 严格权限路由
	strictAuthRouter := r.Use(middleware.StrictAuth(jwt, log))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateUser)
	}
	// 非严格权限路由
	noStrictAuthRouter := r.Use(middleware.NoStrictAuth(jwt, log))
	{
		noStrictAuthRouter.POST("/user", userHandler.CreateUser)
	}

	return r
}
