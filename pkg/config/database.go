package config

import (
	"fmt"
	"os"
)

func LoadDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func LoadRedisCredentials() (string, string) {
	return os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD")
}
