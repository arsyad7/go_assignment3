package controllers

import (
	"go_assignment3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	res := models.GetStatus()

	var water models.Water
	var wind models.Wind

	water.Status = res.Water.Status
	water.Message = res.Water.Message
	wind.Message = res.Wind.Message
	wind.Status = res.Water.Status

	c.HTML(http.StatusOK, "template.html", gin.H{
		"water": water,
		"wind":  wind,
	})
}
