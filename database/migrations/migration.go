package migrations

import (
	"belajar-fiber/database"
	"belajar-fiber/model/entity"
)

func RunMigrations() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic("failed to migrate database")
	}

}
