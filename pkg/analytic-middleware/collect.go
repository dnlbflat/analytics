package analytic_middleware

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	Debug string
	DB    *gorm.DB
	Rdb   redis.UniversalClient
}

type Analytic struct {
	db *gorm.DB
}

func New(cfg Config) (Collector, error) {
	return &Analytic{
		db: cfg.DB,
	}, nil
}
