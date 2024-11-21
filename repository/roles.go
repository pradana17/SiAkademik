package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
)

func CreateRole(role *models.Role) error {
	err := database.DB.Create(role).Error

	if err != nil {
		return err
	}
	return nil
}

func GetRoleByID(id uint, role *models.Role) error {
	err := database.DB.Where("id = ?", id).First(role).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteRoleByID(id uint, role *models.Role) error {

	err := database.DB.Where("id=?", id).Delete(role).Error
	if err != nil {
		return err
	}
	return nil
}
