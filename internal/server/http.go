package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/internal/handler"
	"github.com/go-nunu/nunu-layout-advanced/internal/middleware"
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/resp"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	jwt *middleware.JWT,
	userHandler *handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	// 无权限路由
	noAuthRouter := r.Group("/").Use(middleware.RequestLogMiddleware(logger))
	{

		noAuthRouter.GET("/", func(ctx *gin.Context) {
			logger.WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"say": "Hi Nunu!",
			})
		})

		noAuthRouter.POST("/user/register", userHandler.Register)
		noAuthRouter.POST("/user/login", userHandler.Login)
	}
	// 非严格权限路由
	noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// 严格权限路由
	strictAuthRouter := r.Group("/").Use(middleware.StrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return r
}
