package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiEndpoint string
	AppKey      string
	SecretKey   string
}

func Load() *Config {
	// .env 파일 불러오기
	err := godotenv.Load()
	if err != nil {
		log.Println(".env 파일을 찾을 수 없습니다. (환경변수로부터 직접 읽습니다.)")
	}

	return &Config{
		ApiEndpoint: getEnv("API_BASE_URL", ""),
		AppKey:      getEnv("APP_KEY", ""),
		SecretKey:   getEnv("SECRET_KEY", ""),
	}
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
