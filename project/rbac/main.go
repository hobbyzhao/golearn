package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golearn/project/rbac/router"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.RegisterRouter(r)

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080
		fmt.Println("fail to run server", err)
	}
	fmt.Println("Quitting..., start to close process")
}
