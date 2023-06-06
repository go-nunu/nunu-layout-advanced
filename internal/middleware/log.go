package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/md5"
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/uuid"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
	"io"
	"strconv"
	"time"
)

func RequestLogMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 每次请求都初始化一次配置
		trace := md5.Md5(uuid.GenUUID())
		logger.NewContext(ctx, zap.String("trace", trace))
		logger.NewContext(ctx, zap.String("request_method", ctx.Request.Method))
		headers, _ := json.Marshal(ctx.Request.Header)
		logger.NewContext(ctx, zap.String("request_headers", string(headers)))
		logger.NewContext(ctx, zap.String("request_url", ctx.Request.URL.String()))
		if ctx.Request.Body != nil {
			bodyBytes, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
			logger.NewContext(ctx, zap.String("request_params", string(bodyBytes)))
		}
		logger.WithContext(ctx).Info("Request")
		ctx.Next()
	}
}
func ResponseLogMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		startTime := time.Now()
		ctx.Next()
		duration := int(time.Since(startTime).Milliseconds())
		ctx.Header("X-Response-Time", strconv.Itoa(duration))
		logger.WithContext(ctx).Info("Response", zap.Any("response_body", blw.body.String()), zap.Any("time", fmt.Sprintf("%sms", strconv.Itoa(duration))))
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
