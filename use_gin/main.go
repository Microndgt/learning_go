package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func getName()(first string, second string){
	const a, b float32 = 2, 3.0
	fmt.Print(a+b)
	return "may", "june"
}

func main() {
	fn, mn := getName()
	fmt.Print(fn, mn)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}