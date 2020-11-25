package controller

import (
	"github.com/gin-gonic/gin"
	"gocrm/models"
)

func GetProduct(c *gin.Context) {
	code := c.Query("code")
	c.JSON(200, models.GetPro(code))
}
