package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
)

func CreateUserProfile(up *models.UserProfile) error {
	err := database.DB.Create(up).Error

	if err != nil {
		return err
	}
	return nil
}

func UpdateUserProfile(userID uint, up models.UserProfile) error {
	var oldData models.UserProfile

	// Get the existing user data from the database
	get := database.DB.First(&oldData, userID).Error
	if get != nil {
		return get // Return the error if the user is not found
	}

	// Update the user data in the database
	upd := database.DB.Model(&oldData).Updates(up).Error
	if upd != nil {
		return upd // Return the error if the update fails
	}

	return nil // Return nil if both operations are successful
}

func GetUserProfile(userid uint, up *models.UserProfile) error {
	err := database.DB.Where("user_id = ?", userid).First(up).Error
	if err != nil {
		return err
	}
	return nil
}
