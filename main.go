package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "fullstack-postgres"
	port     = 5432
	user     = "admin"
	password = "admin123"
	dbname   = "dev"
)

func main() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s db=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	log.Printf("postgres started at %d port", port)
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", pingHandler)

	r.GET("/cow/:phrase", cowHandler)

	r.GET("/person/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, name)
	})

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
		"title":  "bow to me faithfully",
		"phrase": say,
	})
}
