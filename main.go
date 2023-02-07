package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type form struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

var forms = []form{
	{Title: "Blue Train", Name: "John Coltrane"},
	{Title: "Jeru", Name: "Gerry Mulligan"},
}

func getForm(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, forms)
}

func main() {
	router := gin.Default()
	router.GET("/", getForm)
	router.Run("localhost:8080")
}
