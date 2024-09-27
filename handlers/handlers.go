package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func PostForm(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")

	if name == "" || address == "" {
		c.String(http.StatusBadRequest, "null values not allowed")
		return 
	}
	c.String(http.StatusOK, "Name: %s\nAdress: %s", name, address)
}