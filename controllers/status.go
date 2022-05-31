package controllers

import (
	"go_assignment3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	res := models.GetStatus()

	c.HTML(http.StatusOK, "template.html", gin.H{
		"water": models.Water{
			Status:  res.Water.Status,
			Message: res.Water.Message,
		},
		"wind": models.Wind{
			Status:  res.Wind.Status,
			Message: res.Wind.Message,
		},
	})
}
