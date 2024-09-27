package main

import (
	"github.com/Soyaib10/gin-web-server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default(); // gin router

	r.Static("/static", "./stactic")

	r.GET("/", handlers.HomeHandler)
	r.GET("/form", handlers.GetForm)
	r.POST("/form", handlers.PostForm)

	r.Run(":8080");
}