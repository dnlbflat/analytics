package analytic_middleware

import (
	"analytic-middleware/pkg/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (a *Analytic) Collect(c *gin.Context) {
	var requestBody []byte

	if c.Request.Body != nil {
		requestBody, _ = io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}

	responseBody := &bytes.Buffer{}

	originalWriter := c.Writer
	c.Writer = &bodyWriter{
		ResponseWriter: originalWriter,
		body:           responseBody,
	}

	c.Next()

	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()
	method := c.Request.Method
	path := c.Request.URL.Path

	logEntry := models.LogEntry{
		Method:       method,
		Path:         path,
		StatusCode:   statusCode,
		ClientIP:     clientIP,
		RequestBody:  string(requestBody),
		ResponseBody: responseBody.String(),
	}

	if err := a.db.Create(&logEntry).Error; err != nil {
		c.Error(err)
	}
}
