package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naoki-sawada/golang-rest-example/router"
)

func main() {
	app := gin.Default()
	router.Router(app)
	app.Run(":8080")
}
