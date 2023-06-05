package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout/pkg/helper/md5"
	"github.com/go-nunu/nunu-layout/pkg/helper/uuid"
	"github.com/go-nunu/nunu-layout/pkg/log"
	"go.uber.org/zap"
	"io"
	"strconv"
	"strings"
	"time"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
func RequestLogMiddleware(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每次请求都初始化一次配置
		trace := md5.Md5(uuid.GenUUID())
		log.NewContext(ctx, zap.String("trace", trace))
		log.NewContext(ctx, zap.String("request_method", ctx.Request.Method))
		headers, _ := json.Marshal(ctx.Request.Header)
		log.NewContext(ctx, zap.String("request_headers", string(headers)))
		log.NewContext(ctx, zap.String("request_url", ctx.Request.URL.String()))
		if ctx.Request.Body != nil {
			bodyBytes, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
			log.NewContext(ctx, zap.String("request_params", string(bodyBytes)))
		}
		ctx.Next()
	}
}
func ResponseLogMiddleware(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		startTime := time.Now()
		ctx.Next()
		duration := int(time.Since(startTime).Milliseconds())
		ctx.Header("X-Response-Time", strconv.Itoa(duration))
		if ctx.Request.URL.Path == "/cos/object" && ctx.Request.Method == "POST" {
			return
		}
		if strings.Contains(ctx.Request.URL.Path, "storage") {
			return
		}
		log.WithContext(ctx).Info("响应返回", zap.Any("response_body", blw.body.String()), zap.Any("time", fmt.Sprintf("%sms", strconv.Itoa(duration))))
		statusCode := ctx.Writer.Status()
		fmt.Println(statusCode)
		//if statusCode >= 400 {
		//ok this is an request with error, let's make a record for it
		// now print body (or log in your preferred way)
		//fmt.Println("Response body: " + blw.body.String())
		//}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
