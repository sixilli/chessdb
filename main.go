package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sixilli/chess-api/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/games", controllers.GetGame)
	r.Run(":8080")
}
