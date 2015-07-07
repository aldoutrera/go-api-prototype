package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)

type Tour struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Site struct {
	Title string
}

func findTours(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT id, name FROM tours")
	if err != nil {
	}
	return rows, err
}

func main() {
	// Connect to the db
	db, err := sql.Open("postgres", "user=desarrollo dbname=toori sslmode=disable")
	if err != nil {
	}
	defer db.Close()

	// Creates a gin router + logger and recovery (crash-free) middlewares
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		site := Site{ Title: "title of the site" }
		c.HTML(http.StatusOK, "index.tmpl", site)
	})

	r.GET("/tours", func(c *gin.Context) {
		tours, err := findTours(db)
		if err != nil {
			fmt.Println(err)
		}
		defer tours.Close()
		var results []Tour
		for tours.Next() {
			tour := Tour{}
			err := tours.Scan(&tour.Id, &tour.Name)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, tour)
		}
		json_object := map[string][]Tour{ "tours": results }
		c.JSON(http.StatusOK, json_object)
	})

	r.GET("/pong", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
