package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
)

func GetUserByUsername(username string, user *models.User) error {
	err := database.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(id uint, user *models.User) error {
	err := database.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(user *models.User) error {
	err := database.DB.Create(user).Error

	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(userID uint, user models.User) error {
	var oldData models.User

	// Get user yang ada di DB
	get := database.DB.First(&oldData, userID).Error
	if get != nil {
		return get // Return error jika user tidak ditemukan
	}

	// Update user data di database
	upd := database.DB.Model(&oldData).Updates(user).Error
	if upd != nil {
		return upd // Return error jika update gagal
	}

	return nil // return jika update success
}
