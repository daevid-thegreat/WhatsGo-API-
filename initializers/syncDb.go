package initializers

import "whatsgo/models"

func SyncDb() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
