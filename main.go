package main

import (
	"github.com/gin-gonic/gin"
	"gocrm/route"
	//"io"
	//"os"
)

func main() {
	//gin日志
	//f, _ := os.Create("/tmp/gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	v1 := r.Group("/api/v1.0/internal")
	{
		v1.GET("/status", route.Heartbeat)
		v1.GET("/configs", route.GetConfigs)
		v1.POST("/configs", route.PostConfig)
		v1.POST("/all-configs", route.AllConfig)
		v1.GET("/oplog", route.GetOplog)
		v1.GET("/tasks", route.GetTask)
		v1.POST("/tasks", route.PostTask)
		v1.DELETE("/tasks", route.DeleteTask)
		v1.GET("/pro", route.GetProduct)
	}

	r.Run(":9999")
}
