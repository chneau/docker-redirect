package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	redirectTo := os.Getenv("REDIRECT_TO")
	println("REDIRECT_TO:", redirectTo)
	r := gin.Default()
	r.Any("/*_", func(c *gin.Context) {
		c.Redirect(307, redirectTo)
	})
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
