package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
	"fmt"
)

func GetUserProfile(userID uint) (*models.UserProfile, error) {
	var up models.UserProfile
	err := repository.GetUserProfile(userID, &up)
	if err != nil {
		return nil, errors.New("error retrieving userProfile")
	}
	return &up, nil
}

func CreateUserProfile(up *models.UserProfile) error {
	// Memanggil fungsi repository untuk menyimpan data user
	return repository.CreateUserProfile(up)
}

func UpdateUserProfile(userID uint, up models.UserProfile) error {
	existingUser, err := GetUserProfile(userID)
	if err != nil {
		return errors.New("user profile not found11")
	}

	fmt.Printf("%d", existingUser.UserID)

	fmt.Printf("%d", up.UserID)
	// Logika untuk mencegah perubahan role
	// if up.UserID != existingUser.UserID {
	// 	return errors.New("user id cannot be changed")
	// }

	up.UserID = existingUser.UserID

	// Panggil repository untuk mengupdate user
	err = repository.UpdateUserProfile(existingUser.ID, up)
	if err != nil {
		return errors.New("failed to update user")
	}
	return nil
}
