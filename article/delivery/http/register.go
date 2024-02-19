package http

import (
	"clean-arch-demo/article"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc article.UseCase) {
	h := NewHandler(uc)

	articles := router.Group("/articles")
	{
		articles.POST("", h.CreateArticle)
		articles.GET("", h.GetArticle)
		articles.DELETE("", h.DeleteArticle)
	}
}
