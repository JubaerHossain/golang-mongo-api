package services

import "github.com/JubaerHossain/golang-mongodb-api/models"

type PostService interface {
	CreatePost(*models.CreatePostRequest) (*models.DBPost, error)
	UpdatePost(string, *models.UpdatePost) (*models.DBPost, error)
	FindPostById(string) (*models.DBPost, error)
	FindPosts(page int, limit int) ([]*models.DBPost, error)
	DeletePost(string) error
}
