package blog

import (
	"go-blog/internal/blog/controller"
	"go-blog/internal/blog/repository"
	"go-blog/internal/blog/service"

	"github.com/gin-gonic/gin"
)

func InitializeBlogModule(r *gin.Engine) {
	// Initialize repository, service, and controller
	blogRepo := &repository.BlogRepository{}
	blogService := &service.BlogService{Repository: blogRepo}
	blogController := &controller.BlogController{Service: blogService}

	// Define blog routes
	r.GET("/blogs", blogController.GetBlogs)
}
