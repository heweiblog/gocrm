package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocrm/route"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"unique"`
	Price uint
}

func main() {
	dsn := "root:123456@tcp(192.168.5.41:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, "code = ?", "D42")
	fmt.Println(product)

	//gin日志
	f, _ := os.Create("/tmp/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	v1 := r.Group("/api/v1.0/internal")
	{
		v1.GET("/status", route.Heartbeat)
		v1.GET("/configs", route.GetConfig)
		v1.POST("/configs", route.PostConfig)
		v1.POST("/all-configs", route.AllConfig)
		v1.GET("/oplog", route.GetOplog)
		v1.GET("/tasks", route.GetTask)
		v1.POST("/tasks", route.PostTask)
		v1.DELETE("/tasks", route.DeleteTask)
	}

	r.Run(":9999")
}
