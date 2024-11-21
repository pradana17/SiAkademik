package controllers

import (
	"SiAkademik/services"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Simpan informasi request
		method := c.Request.Method
		endpoint := c.Request.URL.Path

		userID, ok := c.Get("userid")
		if !ok {
			// Jika userID tidak ditemukan, gunakan default (misal 0)
			userID = uint(0)
		}

		// Lakukan type assertion untuk memastikan tipe uint64
		userIDUint, ok := userID.(uint)
		if !ok {
			// Jika gagal konversi, gunakan default (misal 0)
			userIDUint = uint(0)
		}

		// Baca body request
		var requestBody map[string]interface{}
		if c.Request.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset body untuk dibaca ulang di handler
			_ = json.Unmarshal(bodyBytes, &requestBody)
		}

		// Lanjutkan ke handler berikutnya
		c.Next()

		// Tangkap response status dan durasi
		statusCode := c.Writer.Status()
		duration := time.Since(startTime).Seconds()

		// Buat log response (dapat ditingkatkan jika ingin menangkap body response)
		responseBody := gin.H{
			"status": statusCode,
		}

		// Simpan log audit ke database melalui service
		services.CreateAuditLog(
			userIDUint, // userID (sesuaikan dengan autentikasi jika ada)
			method,
			endpoint,
			requestBody,
			responseBody,
			statusCode,
			duration,
		)
	}
}
