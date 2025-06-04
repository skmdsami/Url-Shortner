package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"log"
	_ "log"

	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func main() {
	r := gin.Default()
	dsn := "user=postgres password=phani dbname=url_shortner sslmode=disable"
	var err error
	DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	r.POST("/", func(c *gin.Context) {
		url := c.PostForm("url")
		log.Println("url:", url)
		var hashID string
		err := DB.QueryRow("INSERT INTO urls (url) VALUES ($1) RETURNING id", url).Scan(&hashID)
		if err != nil {
			log.Fatal("Error inserting url")
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		c.JSON(200, gin.H{
			"url": "http://localhost:8080/" + hashID,
		})
	})

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var url string
		err := DB.QueryRow("SELECT url FROM urls WHERE id = $1", id).Scan(&url)

		if err != nil {
			log.Fatal("Error while fetching url")
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		var count int
		err1 := DB.QueryRow("UPDATE urls SET count = count + 1 WHERE id = $1 RETURNING count", id).Scan(&count)
		if err1 != nil {
			log.Fatal("Error while updating count")
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		log.Println("count was ", count)
		c.Redirect(302, url)
	})

	r.GET("/count/:id", func(c *gin.Context) {
		id := c.Param("id")
		var count int
		err := DB.QueryRow("SELECT count FROM urls WHERE id = $1", id).Scan(&count)
		if err != nil {
			log.Fatal("Error while getting count")
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		c.JSON(200, gin.H{
			"count": count,
		})
	})

	r.Run(":8080")
}

/*

route-1 : from the hash url/hash_id redirecting to the actual website
route-2 : sending the actual url via post request and getting the short url
route-3 : getting the number of people clicked on that link

*/
