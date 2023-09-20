package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onekeyinfo/internal/app/util"
)

func Root(r *gin.Engine) {
	r.StaticFile("/favicon.ico", "web/favicon.ico")
	content, _ := util.GetContent("web/index.html")
	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html;charset=utf-8", []byte(content))
	})
}
