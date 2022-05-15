package config

import (
	// "fmt"
	// "indah_cantik/model"

	"indah_cantik/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Config struct {
// 	DB_Username string
// 	DB_Password string
// 	DB_Port     string
// 	DB_Host     string
// 	DB_Name     string
// }

// func Connect() *gorm.DB {
// 	config := Config{
// 		DB_Username: "root",
// 		DB_Password: "",
// 		DB_Port:     "3306",
// 		DB_Host:     "localhost",
// 		DB_Name:     "tugas",
// 	}

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		config.DB_Username,
// 		config.DB_Password,
// 		config.DB_Host,
// 		config.DB_Port,
// 		config.DB_Name,
// 	)

// 	var err error
// 	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	return DB
// }

var DB *gorm.DB

var JwtSecret string

func InitDB() {
	connectionString := "root:@/tugas?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func InitMigrate() {
	InitDB()
	DB.AutoMigrate(&model.User{}, &model.Product{})
}
