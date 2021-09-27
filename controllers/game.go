package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sixilli/chess-api/externalapis/chessdotcom"
)

func GetGame(c *gin.Context) {
	games, err := chessdotcom.GetPlayerGames("sixilli", "2021", "08")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"games": games,
	})
}
