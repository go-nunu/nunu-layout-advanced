package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/resp"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"regexp"
)

type JWT struct {
	key string
}
type MyCustomClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

// NewJwt https://pkg.go.dev/github.com/golang-jwt/jwt/v5
func NewJwt(conf *viper.Viper) *JWT {
	return &JWT{key: conf.GetString("security.jwt.key")}
}
func (j *JWT) GenToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, MyCustomClaims{
		UserId: 1,
	})

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := token.SignedString(j.key)

	fmt.Println(tokenString, err)
	return tokenString

}
func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	re, _ := regexp.Compile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func NoAuth(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.WithContext(ctx).Info("建立请求")
		ctx.Next()
	}
}

// StrictAuth 严格权限
func StrictAuth(j *JWT, log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			log.WithContext(ctx).Warn("请求未携带token，无权限访问", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			resp.HandleError(ctx, http.StatusUnauthorized, 1, "no token", nil)
			ctx.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			log.WithContext(ctx).Error("token error", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			resp.HandleError(ctx, http.StatusUnauthorized, 1, err.Error(), nil)
			ctx.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, log)
		ctx.Next()
	}
}

func NoStrictAuth(j *JWT, log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("accessToken")
		}
		if tokenString == "" {
			tokenString = ctx.Query("accessToken")
		}
		if tokenString == "" {
			log.WithContext(ctx).Info("建立请求")
			ctx.Next()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			log.WithContext(ctx).Info("建立请求")
			ctx.Next()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, log)
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	if ctx.Request.URL.Path == "/cos/object" && ctx.Request.Method == "POST" {
		return
	}
	userInfo := ctx.MustGet("claims").(*MyCustomClaims)
	logger.NewContext(ctx, zap.Int64("UserId", userInfo.UserId))
	logger.WithContext(ctx).Info("建立请求")

	// 统计
}
