package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"errors"
)

func CreateRole(role *models.Role) error {
	// Misalnya, kita bisa menambahkan validasi sebelum membuat user
	if role.Name == "" {
		return errors.New("role name is required")
	}

	// Memanggil fungsi repository untuk menyimpan data user
	return repository.CreateRole(role)
}

func GetRoleByID(roleID uint) (*models.Role, error) {
	var role models.Role
	err := repository.GetRoleByID(roleID, &role)
	if err != nil {
		return nil, errors.New("error retrieving role")
	}
	return &role, nil
}

func DeleteRole(roleID uint) error {
	var role models.Role
	err := repository.DeleteRoleByID(roleID, &role)
	if err != nil {
		return errors.New("error delete role")
	}
	return nil
}
