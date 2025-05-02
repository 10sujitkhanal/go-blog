package main

import (
	"go-blog/internal/blog"

	"github.com/gin-gonic/gin"
)

func main() {

	// db.Connect()

	r := gin.Default()

	blog.InitializeBlogModule(r)

	r.Run(":8080")
}
