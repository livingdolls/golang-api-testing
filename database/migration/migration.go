package migration

import (
	"fmt"
	"gofiber/database"
	"gofiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{});

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Migrated")
}