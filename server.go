package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	server := flag.Bool("server", false, "Whether this is Tig server.")
	flag.Parse()
	if *server {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run()
	} else {
		subCmd := os.Args[1]
		switch subCmd {
		case "submit":
			fmt.Println("submit command")
		case "export":
			fmt.Println("export command")
		}
	}
}
