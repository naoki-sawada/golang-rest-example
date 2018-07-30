package router

import (
	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PutLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func DeleteLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
