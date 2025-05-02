package controller

import (
	"net/http"

	"go-blog/internal/blog/service"

	"github.com/gin-gonic/gin"
)

// BlogController handles HTTP requests for blogs.
type BlogController struct {
	Service *service.BlogService
}

// GetBlogs handles GET /blogs requests.
func (c *BlogController) GetBlogs(ctx *gin.Context) {
	blogs := c.Service.GetBlogs()
	ctx.JSON(http.StatusOK, blogs)
}
