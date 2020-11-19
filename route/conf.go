package route

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gocrm/models"
	"reflect"
)

type Contents struct {
	MsRelease string           `json:"msrelease"`
	Contents  []models.Content `json:"contents"`
}

func GetConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "received",
		"rcode":  0,
		"conf":   "test",
	})
}

type Res struct {
	Responsecode int `json:"responsecode"`
}

type Base map[string]map[string]map[string]func(c *models.Content)

var Hand Base

func Call(c *models.Content) {
	fmt.Println("func recv:", c)
	m := make(map[string]interface{})
	json.Unmarshal(c.Data, &m)
	fmt.Println("func data:", m)
	fmt.Println("func data:", reflect.TypeOf(m["responsecode"]), m["responsecode"])
	code := int(m["responsecode"].(float64))
	fmt.Println("func recv data code:", code)

	var r Res
	json.Unmarshal(c.Data, &r)
	fmt.Println("---func data:", reflect.TypeOf(r), r)
	models.DB.Create(c)
}

func init() {
	Hand = make(Base)
	fmt.Println(Hand)
	m := make(map[string]func(c *models.Content))
	m["responserules"] = Call
	n := make(map[string]map[string]func(c *models.Content))
	n["selfcheck"] = m
	Hand["handle"] = n
	fmt.Println(Hand)
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
			Hand[d.Contents[i].Service][d.Contents[i].Bt][d.Contents[i].Sbt](&d.Contents[i])
		}
		/*
			for _, i := range d.Contents {
				fmt.Println("i:", reflect.TypeOf(i), i)
				fmt.Println("data:", reflect.TypeOf(i.Data), i.Data)
				val, err := i.Data.Value()
				fmt.Println(val, err)
				fmt.Println("val:", reflect.TypeOf(val), val)
				//fmt.Println(i)
				//for k, v := range i.Data {
				//fmt.Println(k, v, reflect.TypeOf(v))
				//}
			}
		*/
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
