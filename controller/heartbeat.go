package controller

import (
	"github.com/gin-gonic/gin"
	"gocrm/models"
	"gocrm/utils"
)

func Heartbeat(c *gin.Context) {
	m := make(map[string]interface{})
	m["status"] = "running"
	m["msrelease"] = utils.MsRelease
	m["devicerelease"] = ""
	m["msversion"], m["deviceversion"] = models.GetOplogId()
	m["softwareversion"] = "gcrm-1.1.1"
	m["licenseinfo"] = "abcABC=="

	c.JSON(200, m)
}
