package main

import (
	"fmt"
	"go/build"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type PingPong struct {
	Message   string `json:"message"`
	NeuesFeld string `json:"neues_feld"`
}

func main() {
	r := gin.Default()
	// r.Static("/public", "./public")
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	// damit der Debugger im Browser verwendet werden kann, werden hier die
	// GO-Quellen als Static-Daten zur Verfügung gestellt
	r.Use(static.Serve("/", static.LocalFile(build.Default.GOPATH+"/src", false)))
	r.Use(static.Serve("/", static.LocalFile(build.Default.GOROOT+"/src", false)))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":    "pong",
			"neues_feld": "Hey",
		})
	})
	r.POST("/save", func(c *gin.Context) {
		var p PingPong

		c.BindJSON(&p)

		c.JSON(200, gin.H{
			"message":    p.Message,
			"neues_feld": p.NeuesFeld,
		})
		fmt.Printf("Message: %s, NeuesFeld: %s\n", p.Message, p.NeuesFeld)
		fmt.Printf("%v\n", p)
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
