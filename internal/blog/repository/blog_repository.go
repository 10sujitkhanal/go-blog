package repository

import models "github.com/10sujitkhanal/go-blog/internal/blog/models"

// BlogRepository handles data fetching logic.
type BlogRepository struct{}

// GetAllBlogs returns all blogs from the data source.
func (r *BlogRepository) GetAllBlogs() []models.Blog {
	return []models.Blog{
		{ID: 1, Title: "First Blog Post", Content: "Content of the first blog.", Author: "John Doe"},
		{ID: 2, Title: "Second Blog Post", Content: "Content of the second blog.", Author: "Jane Smith"},
		{ID: 3, Title: "Go is Awesome!", Content: "Content about Go.", Author: "Alice Johnson"},
	}
}
