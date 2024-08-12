//go:build wireinject
// +build wireinject

package config

import (
	"fmt"
	"github.com/Savitar465/task-manager/app/modules"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbsql = InitDB()
var db = wire.NewSet(dbsql)

func InitApp() *Initialization {
	wire.Build(NewInitialization, db, modules.IssueCtrlSet, modules.IssueServiceSet, modules.IssueRepoSet)
	return nil
}

func InitDB() *gorm.DB {
	// Construct DSN from environment variables
	username := viper.GetString("database.user")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.db")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	// Open database connection
	// Get a database handle.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to database with GORM!")
	return db
}
