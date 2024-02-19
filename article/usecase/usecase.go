package usecase

import (
	"clean-arch-demo/article"
	"clean-arch-demo/models"
	"context"
)

var _ article.UseCase = (*ArticleUseCase)(nil)

type ArticleUseCase struct {
	articleRepo article.Repository
}

func NewArticleUseCase(articleRepo article.Repository) *ArticleUseCase {
	return &ArticleUseCase{
		articleRepo: articleRepo,
	}
}

func (a ArticleUseCase) CreateArticle(ctx context.Context, user *models.User, article *models.Article) error {
	article.UserID = user.ID
	return a.articleRepo.CreateArticle(ctx, article)
}

func (a ArticleUseCase) GetArticle(ctx context.Context, user *models.User) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleUseCase) DeleteArticle(ctx context.Context, user *models.User, id uint) (*models.Article, error) {
	//TODO implement me
	panic("implement me")
}
