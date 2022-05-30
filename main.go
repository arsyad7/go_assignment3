package main

import (
	"go_assignment3/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

const APP_PORT = ":8888"

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("views/template.html")

	r.GET("/random", controllers.GetStatus)

	log.Println("server running at port ", APP_PORT)
	r.Run(APP_PORT)
}
