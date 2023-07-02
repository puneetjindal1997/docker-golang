package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", helloworld)
	r.Run(":8000")
}

func helloworld(c *gin.Context) {
	fmt.Println("hello world")
	c.JSON(http.StatusOK, gin.H{"message": "hello docker"})
}
