package config

import "os"

type Config struct {
	ServerPort string
	DBHost     string
	DBName     string
	DBPassword string
	DBUser     string
	Env        string
}

func LoadConfig() *Config {
	env := getEnv("APP_ENV", "dev")
	var config Config
	if env == "prod" {
		config = Config{
			ServerPort: getEnv("PROD_SERVER_PORT", ":8080"),
			DBHost:     getEnv("PROD_DB_HOST", "prod-db-host"),
			DBUser:     getEnv("PROD_DB_USER", "prod-db-user"),
			DBPassword: getEnv("PROD_DB_PASSWORD", "prod-db-password"),
			DBName:     getEnv("PROD_DB_NAME", "prod-db-name"),
			Env:        env,
		}
	} else {
		config = Config{
			ServerPort: getEnv("DEV_SERVER_PORT", ":8080"),
			DBHost:     getEnv("DEV_DB_HOST", "localhost"),
			DBUser:     getEnv("DEV_DB_USER", "dev-user"),
			DBPassword: getEnv("DEV_DB_PASSWORD", "dev-password"),
			DBName:     getEnv("DEV_DB_NAME", "dev-db-name"),
			Env:        env,
		}
	}
	return &config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
