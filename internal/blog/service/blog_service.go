package service

import (
	models "github.com/10sujitkhanal/go-blog/internal/blog/models"

	"github.com/10sujitkhanal/go-blog/internal/blog/repository"
)

// BlogService handles the business logic for blogs.
type BlogService struct {
	Repository *repository.BlogRepository
}

// GetBlogs returns a list of blogs.
func (s *BlogService) GetBlogs() []models.Blog {
	return s.Repository.GetAllBlogs()
}
