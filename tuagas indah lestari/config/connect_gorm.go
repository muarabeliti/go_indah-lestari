package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db *gorm.DB
	// mainDB *gorm.DB
)

func SetEnv() {
	os.Setenv("PORT", "3000")

	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_DATABASE", "golang")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")

	os.Setenv("SENTRY", "https://examplePublicKey@o0.ingest.sentry.io/0")
	os.Setenv("ENV", "development")
	os.Setenv("RELEASE", "golangapp@0.0.1")
}

func GetDB() *gorm.DB {
	return db
}

func ConnectGorm() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	fmt.Println(dsn)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         initLog(),
		NamingStrategy: initNamingStrategy(),
	})

	if err != nil {
		fmt.Println(err)
		panic("Fail To Connect Database")
	}

	sql, _ := db.DB()
	sql.SetMaxOpenConns(120)
	sql.SetMaxIdleConns(20)
}

func initLog() logger.Interface {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		Colorful:      true,
		LogLevel:      logger.Info,
		SlowThreshold: time.Second,
	})
	return newLogger
}

func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: false,
		TablePrefix:   "",
	}
}
