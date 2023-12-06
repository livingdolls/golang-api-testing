package database

import (
	"gofiber/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error;
	const MYSQL = "yurina:hirate@tcp(127.0.0.1:3306)/gofiber?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL;
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{});

	if err != nil {
		panic("Cannot connect to database")
	}

	db.AutoMigrate(&entity.User{})

	DB = db
}