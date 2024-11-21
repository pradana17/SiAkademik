package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
)

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := repository.GetUserByID(userID, &user)
	if err != nil {
		return nil, errors.New("error retrieving role")
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	// validasi sebelum membuat user
	if user.Username == "" {
		return errors.New("username is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	// Memanggil fungsi repository untuk menyimpan data user
	return repository.CreateUser(user)
}

// UpdateUser untuk melakukan update username password
func UpdateUser(userID uint, user models.User) error {
	existingUser, err := GetUserByID(userID)
	if err != nil {
		return errors.New("user not found11")
	}
	if user.RoleID == 0 {
		user.RoleID = existingUser.RoleID
	}
	// Logika untuk mencegah perubahan role
	if user.RoleID != existingUser.RoleID {
		return errors.New("role cannot be changed")
	}
	// Panggil repository untuk mengupdate user
	err = repository.UpdateUser(userID, user)
	if err != nil {
		return errors.New("failed to update user")
	}
	return nil
}
