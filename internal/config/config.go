package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl     string
	JwtSecret string
	Port      string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found,using system env")
	}
	cfg := &Config{
		Port:      getEnv("PORT", "4200"),
		DBUrl:     mustGetEnv("DATABASE_URL"),
		JwtSecret: mustGetEnv("JWT_SECRET"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Missing required env: %s", key)
		return fallback
	}
	return val
}
func mustGetEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Missing required env: %s", key)

	}
	return val
}
