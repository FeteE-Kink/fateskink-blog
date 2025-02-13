package database

import (
	"fmt"
	"log"
	"os"
	"time"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbLogger struct {
	logger.Interface
}

var Db *gorm.DB

func InitDb() *gorm.DB {
	if Db != nil {
		return Db
	}

	var err error
	newLogger := createLogger()

	// Db, err := gorm.Open(mysql.Open(dbConnectionString()), &gorm.Config{
	// 	Logger: newLogger,
	// })
	Db, err = gorm.Open(postgres.Open(dbConnectionString()), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return Db
}

func createLogger() logger.Interface {
	baseLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Don't ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Include params in the SQL log
			Colorful:                  true,        // Enable color
		},
	)

	return &DbLogger{Interface: baseLogger}
}

func dbConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE_NAME"),
		os.Getenv("POSTGRES_SSL_MODE"))
}

func CloseDb() error {
	if Db != nil {
		sqlDb, err := Db.DB()
		if err != nil {
			return err
		}
		log.Println("Closing database connection...")
		return sqlDb.Close()
	}

	return nil
}
