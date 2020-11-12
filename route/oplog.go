package route

import (
	"github.com/gin-gonic/gin"
)

func GetOplog(c *gin.Context) {
	c.JSON(200, gin.H{
		"oplog": "running",
	})
}
