package cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"onekeyinfo/configs"
	"onekeyinfo/internal/app/router"
)

func Execute() {
	appEnv := configs.GetAppEnv()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router.Root(r)
	router.Query(r)
	err := r.Run(appEnv.Port)
	if err != nil {
		log.Println("服务启动失败")
		return
	}
}
