package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Content struct {
	Source   string                 `json:"source"`
	Id       int                    `json:"id"`
	Sservice string                 `json:"service"`
	Bt       string                 `json:"bt"`
	Sbt      string                 `json:"sbt"`
	Op       string                 `json:"op"`
	Data     map[string]interface{} `json:"data"`
}

type Contents struct {
	Contents []Content `json:"contents"`
}

func GetConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "test",
	})
}

func PostConfig(c *gin.Context) {
	var d Contents
	if err := c.BindJSON(&d); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"rcode": 1, "description": "Data format error cannot be handled"})
		return
	} else {
		fmt.Println(d)
		for _, i := range d.Contents {
			fmt.Println(i)
			for k, v := range i.Data {
				fmt.Println(k, v, reflect.TypeOf(v))
			}
		}
	}
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
	})
}

func AllConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "all",
	})
}
