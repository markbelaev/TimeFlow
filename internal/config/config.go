package config

import (
	"flag"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// Структура с токеном
type Config struct {
	Token string
	isDev bool
}

// Функция врзвращающая токен
func Load() *Config {

	// Парсинг флага для DEV бота
	devMode := flag.Bool("dev", false, "enable development mode")
	flag.Parse()

	// Загрузка .env файла
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file")
		os.Exit(1)
	}

	// Выбор токена в зависимости от режима
	var tokenBot string
	if *devMode {
		tokenBot = os.Getenv("TOKEN_BOT_DEV")
		slog.Info("Running in DEV mode")
	} else {
		tokenBot = os.Getenv("TOKEN_BOT")
		slog.Info("Running in PROD mode")
	}

	// Проверка если токен пустой
	if tokenBot == "" {
		slog.Error("Token is empty")
	}

	// Возвращение структуры с токеном
	return &Config{
		Token: tokenBot,
		isDev: *devMode,
	}
}
