package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"fmt"

	"gorm.io/gorm"
)

func AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	// Cari user berdasarkan username
	err := repository.GetUserByUsername(username, &user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("auth user not found")
		}
		return nil, err // Jika ada error lain
	}

	// Verifikasi password
	if user.Password != password {
		return nil, fmt.Errorf("invalid password")
	}

	// Jika berhasil, kembalikan user
	return &user, nil
}
