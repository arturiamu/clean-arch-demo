package http

import (
	"clean-arch-demo/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	useCase auth.UseCase
}

func NewAuthMiddleware(useCase auth.UseCase) gin.HandlerFunc {
	return (&AuthMiddleware{
		useCase: useCase,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.useCase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	c.Set(auth.CtxUserKey, user)
}
