package controller

import (
	"github.com/gin-gonic/gin"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Heartbeat(c *gin.Context) {
	m := make(map[string]interface{})
	m["status"] = "running"
	m["msrelease"] = "1945982123"
	m["msversion"] = 0
	m["devicerelease"] = 0
	m["softwareversion"] = "gcrm-1.1.1"
	m["licenseinfo"] = "abcABC=="
	//m["test"] = People{Name: "hw", Age: 18}
	//m["test"] = struct{Name:string
	//Age int
	//}{"hw", 18}

	c.JSON(200, m)
}
