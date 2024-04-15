package router

import (
	"log"
	"net/http"
	"onekeyinfo/internal/app/util"

	"github.com/gin-gonic/gin"
)

func Root(r *gin.Engine) {
	r.StaticFile("/favicon.ico", "web/favicon.ico")
	content, err := util.GetContent("web/index.html")
	if err != nil {
		log.Print("网站加载失败.", err)
	}
	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html;charset=utf-8", []byte(content))
	})
}
