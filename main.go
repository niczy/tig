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
			// Submit latest patch to server.
			fmt.Println("submit command")
		case "export":
			// Example local change and master, generate patch file.
			fmt.Println("export command")
		}
	}
}
