package server

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/go-nunu/nunu-layout-advanced/api/v1"
	"github.com/go-nunu/nunu-layout-advanced/docs"
	"github.com/go-nunu/nunu-layout-advanced/internal/handler"
	"github.com/go-nunu/nunu-layout-advanced/internal/middleware"
	"github.com/go-nunu/nunu-layout-advanced/pkg/jwt"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/go-nunu/nunu-layout-advanced/pkg/server/http"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler handler.UserHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})

	v1 := s.Group("/v1")
	// No route group has permission
	noAuthRouter := v1
	{
		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}
	// Non-strict permission routing group
	noStrictAuthRouter := v1.Use(middleware.NoStrictAuth(jwt, logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// Strict permission routing group
	strictAuthRouter := v1.Use(middleware.StrictAuth(jwt, logger))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return s
}
