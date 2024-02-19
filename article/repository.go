package article

import (
	"clean-arch-demo/models"
	"context"
)

type Repository interface {
	CreateArticle(ctx context.Context, article *models.Article) error
	GetArticle(ctx context.Context, id uint) (*models.Article, error)
	DeleteArticle(ctx context.Context, id uint) error
}
