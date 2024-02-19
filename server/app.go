package server

import (
	"clean-arch-demo/article"
	articleHttp "clean-arch-demo/article/delivery/http"
	articleLocalstorage "clean-arch-demo/article/repository/localstorage"
	articleUseCase "clean-arch-demo/article/usecase"
	"clean-arch-demo/auth"
	authHttp "clean-arch-demo/auth/delivery/http"
	authLocalstorage "clean-arch-demo/auth/repository/localstorage"
	authUseCase "clean-arch-demo/auth/usecase"
	"clean-arch-demo/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	authUseCase    auth.UseCase
	articleUseCase article.UseCase
}

func NewApp() *App {
	//db := initDB()

	userRepo := authLocalstorage.NewUserLocalStorage()
	articleRepo := articleLocalstorage.NewArticleLocalStorage()

	return &App{
		authUseCase: authUseCase.NewAuthUseCase(
			userRepo,
			config.DefaultConf.Auth.HashSalt,
			[]byte(config.DefaultConf.Auth.SigningKey),
			time.Duration(config.DefaultConf.Auth.TokenTTL)),

		articleUseCase: articleUseCase.NewArticleUseCase(articleRepo),
	}
}

func (a *App) Run(port int) error {
	router := gin.Default()

	authHttp.RegisterHTTPEndpoints(router, a.authUseCase)
	authMiddleware := authHttp.NewAuthMiddleware(a.authUseCase)

	api := router.Group("/api", authMiddleware)

	articleHttp.RegisterHTTPEndpoints(api, a.articleUseCase)

	a.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *gorm.DB {
	return nil
}
