package repository

import (
	"SiAkademik/database"
	"SiAkademik/models"
)

func SaveAuditLog(auditLog models.Auditlog) error {
	return database.DB.Create(&auditLog).Error
}
