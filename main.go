package main

import (
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", pingHandler)

	r.GET("/cow/:phrase", cowHandler)

	r.Run(":8080")
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func cowHandler(c *gin.Context) {
	phrase := c.Param("phrase")
	say, _ := cowsay.Say(cowsay.Phrase(phrase), cowsay.Type("beavis.zen"))
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "bow to me splendidly",
		"phrase": say,
	})
}
