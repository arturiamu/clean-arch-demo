package localstorage

import (
	"clean-arch-demo/article"
	"clean-arch-demo/auth"
	"clean-arch-demo/models"
	"context"
	"sync"
)

var _ article.Repository = (*ArticleLocalStorage)(nil)

type ArticleLocalStorage struct {
	articles map[uint]*models.Article
	mutex    *sync.Mutex
}

func NewArticleLocalStorage() *ArticleLocalStorage {
	return &ArticleLocalStorage{
		articles: make(map[uint]*models.Article),
		mutex:    new(sync.Mutex),
	}
}

func (a *ArticleLocalStorage) CreateArticle(ctx context.Context, article *models.Article) error {
	a.mutex.Lock()
	a.articles[article.ID] = article
	a.mutex.Unlock()

	return nil
}

func (a *ArticleLocalStorage) GetArticle(ctx context.Context, id uint) (*models.Article, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	for _, a := range a.articles {
		if a.ID == id {
			return a, nil
		}
	}

	return nil, auth.ErrUserNotFound
}

func (a *ArticleLocalStorage) DeleteArticle(ctx context.Context, id uint) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	delete(a.articles, id)
	return nil
}
