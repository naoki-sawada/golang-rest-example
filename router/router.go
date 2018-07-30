package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-sawada/golang-rest-example/middleware"
)

func Router(route *gin.Engine) {
	authorized := route.Group("/")

	authorized.Use(middleware.Authorization())
	{
		authorized.GET("/login", GetLogin)
		authorized.POST("/login", PostLogin)
		authorized.PUT("/login/:id", PutLogin)
		authorized.DELETE("/login/:id", DeleteLogin)
	}
}
