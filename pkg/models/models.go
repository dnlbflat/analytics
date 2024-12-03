package models

type LogEntry struct {
	ID           uint `gorm:"primaryKey"`
	Method       string
	Path         string
	StatusCode   int
	ClientIP     string
	RequestBody  string
	ResponseBody string
}
