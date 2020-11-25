package route

import (
	"github.com/gin-gonic/gin"
	"gocrm/controller"
	//"io"
	//"os"
)

func Server() {
	//gin日志
	//f, _ := os.Create("/tmp/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	v1 := r.Group("/api/v1.0/internal")
	{
		v1.GET("/status", controller.Heartbeat)
		v1.GET("/configs", controller.GetConfigs)
		v1.POST("/configs", controller.PostConfig)
		v1.POST("/all-configs", controller.AllConfig)
		v1.GET("/oplog", controller.GetOplog)
		v1.GET("/tasks", controller.GetTask)
		v1.POST("/tasks", controller.PostTask)
		v1.DELETE("/tasks", controller.DeleteTask)
		v1.GET("/pro", controller.GetProduct)
	}

	r.Run(":9999")
}
