package server

import (
	"context"
	"github.com/Savitar465/task-manager/config"
	_ "github.com/Savitar465/task-manager/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(port string) error {
	// InitViper gin handler

	init := config.InitApp()
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	api := router.Group("/api/v1")
	config.RegisterHTTPEndpoints(api, init)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
