package main

import (
	"context"
	"log"
	controller "main/controllers"
	"main/ent"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func main() {

	databaseURL := "root:@tcp(localhost:3306)/entGo?parseTime=True"
	client, err := ent.Open("mysql", databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}

	client.Debug()
	defer client.Close()

	log.Println("Connected to db")

	// Automatically migrate your schema
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	route := gin.Default()

	// middle ware test start
	route.Use(Logger())
	route.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})
	route.POST("/json", controller.GetJSON())
	// middle ware test end

	route.POST("/user", controller.CreateUser(client))
	route.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
