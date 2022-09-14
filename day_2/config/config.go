package config

import (
	"day_2/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialMigration() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		"root",
		"rootpassword",
		"127.0.0.1",
		"3307",
		"agmc",
		"utf8mb4",
		"true",
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
