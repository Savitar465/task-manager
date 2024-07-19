package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	//db := initDB()

	//userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	//bookmarkRepo := bmmongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))

	return &App{
		//bookmarkUC: bmusecase.NewBookmarkUseCase(bookmarkRepo),
		//authUC: authusecase.NewAuthUseCase(
		//	userRepo,
		//	viper.GetString("auth.hash_salt"),
		//	[]byte(viper.GetString("auth.signing_key")),
		//	viper.GetDuration("auth.token_ttl"),
		//),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	a.httpServer = &http.Server{
		Addr:           ":" + port,
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
