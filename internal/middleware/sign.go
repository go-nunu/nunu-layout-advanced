package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/md5"
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/resp"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/spf13/viper"
	"net/http"
	"sort"
	"strings"
)

func SignMiddleware(logger *log.Logger, conf *viper.Viper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timestamp, ok := ctx.Request.Header["Timestamp"]
		if !ok || len(timestamp) == 0 {
			resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
			ctx.Abort()
			return
		}
		nonce, ok := ctx.Request.Header["Nonce"]
		if !ok || len(nonce) == 0 {
			resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
			ctx.Abort()
			return
		}
		sign, ok := ctx.Request.Header["Sign"]
		if !ok || len(sign) == 0 {
			resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
			ctx.Abort()
			return
		}
		appVersion, ok := ctx.Request.Header["App-Version"]
		if !ok || len(appVersion) == 0 {
			resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
			ctx.Abort()
			return
		}

		data := map[string]string{}
		data["AppKey"] = conf.GetString("security.api_sign.app_key")
		data["Timestamp"] = timestamp[0]
		data["Nonce"] = nonce[0]
		data["AppVersion"] = appVersion[0]

		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
		//拼接
		str := ""
		for _, k := range keys {
			str += k + data[k]
		}
		str += conf.GetString("security.api_sign.app_security")

		if sign[0] != strings.ToUpper(md5.Md5(str)) {
			resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
