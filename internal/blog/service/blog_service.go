package service

import (
	models "go-blog/internal/blog/models"

	"go-blog/internal/blog/repository"
)

// BlogService handles the business logic for blogs.
type BlogService struct {
	Repository *repository.BlogRepository
}

// GetBlogs returns a list of blogs.
func (s *BlogService) GetBlogs() []models.Blog {
	return s.Repository.GetAllBlogs()
}
