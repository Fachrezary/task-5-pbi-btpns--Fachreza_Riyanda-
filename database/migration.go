// database/migration.go

package database

import (
	"projectapi/app"
)

// MigrateDB migrates the database schema.
func MigrateDB() {
	DB.AutoMigrate(&app.User{})
	DB.AutoMigrate(&app.Photo{})

	// Add foreign key constraints or other specific migrations here
	DB.Model(&app.Photo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}