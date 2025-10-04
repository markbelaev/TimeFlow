package main

import (
	"log/slog"
	"os"
	"timeflow/internal/app"
	"timeflow/internal/config"
	"timeflow/internal/data"
	"timeflow/internal/handlers"
)

func main() {

	// Показ тега
	data.Tag()

	// Вызов логера
	app.SetupLooger()

	// Запуск API
	router := app.SetupRouter()

	go func() {
		slog.Info("Starting API...")
		router.Run(":8080")
	}()

	// Загрузка конфига
	cfg := config.Load()

	// Создание бота
	bot, err := app.SetupBot(cfg)
	if err != nil {
		slog.Error("Error creating bot", "error", err)
		os.Exit(1)
	}

	// Регистрация всех обработчиков
	handlers.RegisterAll(bot)

	// Запуск бота
	slog.Info("Starting bot...")
	bot.Start()
}
