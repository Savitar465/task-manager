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

//func InitApp() *config.Initialization {
//	db := initDB()
//	wire.NewSet(db)
//	wire.Build(config.NewInitialization, db, modules.IssueCtrlSet, modules.IssueServiceSet, modules.IssueRepoSet)
//	return nil
//}

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

//func InitDB() *gorm.DB {
//	// Construct DSN from environment variables
//	username := viper.GetString("database.user")
//	password := viper.GetString("database.password")
//	host := viper.GetString("database.host")
//	port := viper.GetString("database.port")
//	database := viper.GetString("database.db")
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
//
//	// Open database connection
//	// Get a database handle.
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %v", err)
//	}
//
//	sqlDB, err := db.DB()
//	if err != nil {
//		log.Fatalf("Failed to get database connection: %v", err)
//	}
//
//	err = sqlDB.Ping()
//	if err != nil {
//		log.Fatalf("Failed to ping database: %v", err)
//	}
//
//	fmt.Println("Connected to database with GORM!")
//	return db
//}
