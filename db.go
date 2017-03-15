package main

import (
	"github.com/jinzhu/gorm"
)

// MigrateDb func
func MigrateDb(db *gorm.DB) {
	if !db.HasTable(&Todo{}) {
		db.CreateTable(&Todo{})
	}

	db.AutoMigrate(&Todo{})
	var count int
	db.Model(&Todo{}).Count(&count)
	if count == 0 {
		db.Create(&Todo{Title: "Create chat", IsComplete: false})
		db.Create(&Todo{Title: "Drink tea", IsComplete: true})
		db.Create(&Todo{Title: "Search job", IsComplete: false})
	}
	// db.DropTable(&Todo{})

}
