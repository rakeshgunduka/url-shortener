package storage

import (
	"fmt"
	"log"
	"os"
	"time"
	"url-shortener-go/config"
	"url-shortener-go/storage/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

// ConnectDB initializes the database connection
func ConnectDB() error {
	user := config.GetConfigValue("DB_USER")
	password := config.GetConfigValue("DB_PASSWORD")
	host := config.GetConfigValue("DB_HOST")
	port := config.GetConfigValue("DB_PORT")
	dbname := config.GetConfigValue("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	migratationErr := DB.AutoMigrate(models.UrlAliasMap{}, models.TokenRange{}, models.Event{})
	if migratationErr != nil {
		panic("failed to migrate database")
	}

	return nil
}
