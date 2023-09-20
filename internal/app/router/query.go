package router

import (
	"github.com/gin-gonic/gin"
	http2 "net/http"
	"onekeyinfo/internal/app/http"
)

func Query(r *gin.Engine) {
	r.GET("/keyInfo", func(ctx *gin.Context) {
		key := ctx.Query("key")
		host := ctx.Query("host")
		info := http.GetKeyInfo(key, host)
		ctx.JSON(http2.StatusOK, gin.H{
			"data":    info,
			"message": "success",
		})
	})
}
