package services

import (
	"SiAkademik/models"
	"SiAkademik/repository"
	"encoding/json"
	"log"
	"time"
)

func CreateAuditLog(userID uint, method string, endpoint string, request interface{}, response interface{}, statusCode int, duration float64) {
	// Konversi request dan response ke JSON string
	reqJSON, err := json.Marshal(request)
	if err != nil {
		log.Printf("Failed to marshal request: %v", err)
		return
	}

	resJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	// Buat data log
	auditLog := models.Auditlog{
		UserID:     &userID,
		Method:     method,
		Endpoint:   endpoint,
		Request:    string(reqJSON),
		Response:   string(resJSON),
		StatusCode: statusCode,
		Duration:   duration,
		CreatedAt:  time.Now(),
	}

	// Simpan log ke database
	if err := repository.SaveAuditLog(auditLog); err != nil {
		log.Printf("Failed to save audit log: %v", err)
	}
}
