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

type Person struct {
	Person       string
	Age          int
	FavoriteFood string
}

func main() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	log.Println(dbinfo)
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
		person := getPerson(name, db)
		c.JSON(http.StatusOK, person)
	})

	r.PUT("/person/create", func(c *gin.Context) {
		var newperson Person
		c.BindJSON(&newperson)
		putPerson(newperson, db)
		c.JSON(http.StatusCreated, newperson)
	})

	r.DELETE("/person/:name", func(c *gin.Context) {
		log.Println("garbage")
		name := c.Param("name")
		result := deletePerson(name, db)
		if result {
			c.Status(http.StatusNoContent)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"person": "not found"})
		}
	})

	r.PATCH("/person/:name/:newName", func(c *gin.Context) {
		name := c.Param("name")
		newName := c.Param("newName")
		patchPersonName(name, newName, db)
		param := getPerson(newName, db)
		c.JSON(http.StatusOK, param)
	})

	r.Run(":8080")
}

func deletePerson(name string, db *sql.DB) bool {
	_, err := db.Exec(`delete from helloworld.person where person.Person = $1`, name)
	if err != nil {
		return false
	}
	return true
}

func putPerson(person Person, db *sql.DB) {
	_, err := db.Exec(`insert into helloworld.person (Person, Age, FavoriteFood) values ($1, $2, $3)`, person.Person, person.Age, person.FavoriteFood)
	if err != nil {
		panic(err)
	}

}

func getPerson(name string, db *sql.DB) (person Person) {
	row := db.QueryRow(`select Age, Person, FavoriteFood from helloworld.person where Person =$1`, name)
	err := row.Scan(&person.Age, &person.Person, &person.FavoriteFood)
	switch err {
	case sql.ErrNoRows:
		return
	case nil:
		return
	default:
		panic(err)
	}
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
func patchPersonName(name string, newName string, db *sql.DB) {
	_, err := db.Exec(`update helloworld.person set Person = $1 where Person = $2`, newName, name)
	if err != nil {
		panic(err)
	}
}
