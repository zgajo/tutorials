// models/setup.go

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SetupModels: allows us to create a connection with our database and migrate our model’s schema.
func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Ok je")
	}

	db.AutoMigrate(&Book{})

	return db
}
