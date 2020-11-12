package route

import (
	"github.com/gin-gonic/gin"
)

func Heartbeat(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":          "running",
		"msrelease":       "2020",
		"msversion":       0,
		"devicerelease":   0,
		"softwareversion": "gocrm-1.0.0",
		"licenseinfo":     "abcABC==",
	})
}
