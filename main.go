package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PingPong struct {
	Message   string `json: "message"`
	NeuesFeld string `json: "neuesfeld"`
}

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "pong",
			"neuesfeld": "Hey",
		})
	})
	r.POST("/save", func(c *gin.Context) {
		var p PingPong
		// scheinbar funktioniert BindJSON mit Feldname mit Underscore nicht
		// z.B. neues_feld
		c.BindJSON(&p)

		c.JSON(200, gin.H{
			"message":   p.Message,
			"neuesfeld": p.NeuesFeld,
		})
		fmt.Printf("message: %s, neues_feld: %s\n", p.Message, p.NeuesFeld)
		fmt.Printf("%v\n", p)
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
