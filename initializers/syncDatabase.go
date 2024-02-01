package initializers

import "auth-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
