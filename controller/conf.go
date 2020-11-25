package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gocrm/models"
	"gocrm/utils"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Contents struct {
	MsRelease string           `json:"msrelease"`
	Contents  []models.Content `json:"contents"`
}

func testPost(request Contents) {
	url := "http://127.0.0.1:22222/"

	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(request)

	fmt.Println(requestBody)

	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(req)
	client := &http.Client{}
	fmt.Println(client)
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func PostConfig(c *gin.Context) {
	var d Contents
	if err := c.BindJSON(&d); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"rcode": 1, "description": "Data format error cannot be handled"})
		return
	} else {
		fmt.Println("recv:", reflect.TypeOf(d), d)
		for i := 0; i < len(d.Contents); i++ {
			key := d.Contents[i].Service + d.Contents[i].Bt + d.Contents[i].Sbt
			if _, ok := utils.CheckMethods[key]; ok {
				//校验模块
				if res := utils.CheckMethods[key](&d.Contents[i]); res == "" {
					//校验通过 入队(通道)
					testPost(d)
					fmt.Println("数据校验通过")
				} else {
					//校验失败，直接记录oplog，或者发到通道，在另一个协程中收取写入
					fmt.Println("数据校验失败", res)
				}
			} else {
				//service bt sbt或对应的函数有问题，直接记录oplog失败，或者发到通道，在另一个协程中收取写入
				fmt.Println("数据大格式错误")
			}
		}
	}
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
	})
}

func GetConfigs(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "all",
	})
}

func AllConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "all",
	})
}
