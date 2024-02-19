package http

import (
	"clean-arch-demo/article"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase article.UseCase
}

func NewHandler(useCase article.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateArticle(c *gin.Context) {
	h.useCase.CreateArticle(c.Request.Context(), nil, nil)
}

func (h *Handler) GetArticle(c *gin.Context) {
	h.useCase.CreateArticle(c.Request.Context(), nil, nil)
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	h.useCase.CreateArticle(c.Request.Context(), nil, nil)
}
