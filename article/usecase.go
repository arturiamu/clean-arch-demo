package article

import (
	"clean-arch-demo/models"
	"context"
)

type UseCase interface {
	CreateArticle(ctx context.Context, user *models.User, article *models.Article) error
	GetArticle(ctx context.Context, user *models.User) (*models.Article, error)
	DeleteArticle(ctx context.Context, user *models.User, id uint) (*models.Article, error)
}
