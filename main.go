package main

import (
	"net/http"

	"github.com/Soyaib10/gin-web-server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default(); // gin router

	r.StaticFS("/static", http.Dir("./stactic"))
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.HomeHandler)
	r.GET("/form", handlers.GetForm)
	r.POST("/form", handlers.PostForm)

	r.Run(":8080");
}