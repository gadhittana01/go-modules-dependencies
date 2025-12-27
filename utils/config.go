package utils

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type BaseConfig struct {
	MigrationURL string
	DBName       string
}

type Config struct {
	DBConnString string
	Port         string
	MigrationURL string
	DBName       string
}

func LoadEnv(paths ...string) {
	if len(paths) > 0 {
		godotenv.Load(paths...)
	} else {
		godotenv.Load()
	}
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ValidateConfig(config *Config) error {
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	return nil
}

func CheckAndSetConfig(configPath, configName string) *Config {
	LoadEnv(configPath + "/" + configName + ".env")

	return &Config{
		DBConnString: GetEnv("DB_CONN_STRING", "postgres://user:password@localhost:5432/parking_lot?sslmode=disable"),
		Port:         GetEnv("PORT", "8000"),
		MigrationURL: GetEnv("MIGRATION_URL", "file://db/migration"),
		DBName:       GetEnv("DB_NAME", "parking_lot"),
	}
}
