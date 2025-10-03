package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// Структура с токеном
type Config struct {
	Token string
}

// Функция врзвращающая токен
func Load() *Config {

	// Загрузка .env файла
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file")
		os.Exit(1)
	}

	// Проверка токена
	tokenBot := os.Getenv("TOKEN_BOT")
	if tokenBot == "" {
		slog.Error("TOKEN_BOT is empty")
		os.Exit(1)
	}

	// Возвращение структуры с токеном
	return &Config{
		Token: tokenBot,
	}
}
