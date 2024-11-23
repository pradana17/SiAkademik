package middlewares

import (
	"SiAkademik/services"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func AuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()

		// Simpan informasi request
		method := c.Request.Method
		endpoint := c.Request.URL.Path

		userID, ok := c.Get("userid")
		if !ok {
			userID = uint(0)
		}

		// // Lakukan type assertion untuk memastikan tipe uint64
		userIDUint, ok := userID.(uint)
		if !ok {
			userIDUint = uint(0)
		}

		// Baca body request
		var requestBody map[string]interface{}
		if c.Request.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset body untuk dibaca ulang di handler
			_ = json.Unmarshal(bodyBytes, &requestBody)
		}
		responseBodyBuffer := &bytes.Buffer{}
		writerWrapper := &responseWriterWrapper{
			ResponseWriter: c.Writer,
			body:           responseBodyBuffer,
		}
		c.Writer = writerWrapper
		// Lanjutkan ke handler berikutnya
		c.Next()

		// Tangkap response status dan durasi
		statusCode := c.Writer.Status()
		duration := time.Since(startTime).Seconds()

		var responseBody map[string]interface{}
		_ = json.Unmarshal(writerWrapper.body.Bytes(), &responseBody)

		// Simpan log audit ke database melalui service
		services.CreateAuditLog(
			userIDUint, // userID
			method,
			endpoint,
			requestBody,
			responseBody,
			statusCode,
			duration,
		)
	}

}

func (r *responseWriterWrapper) Write(b []byte) (int, error) {
	r.body.Write(b) // Salin data yang ditulis ke buffer
	return r.ResponseWriter.Write(b)
}

type responseWriterWrapper struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
