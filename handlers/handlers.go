package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.File("./static/index.html")
}

func GetForm(c *gin.Context) {
	c.File("./static/form.html")
}

func PostForm(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")

	if name == "" || address == "" {
		c.String(http.StatusBadRequest, "null values not allowed")
		return 
	}
	c.String(http.StatusOK, fmt.Sprintf("Name: %s\nAdress: %s", name, address))
}