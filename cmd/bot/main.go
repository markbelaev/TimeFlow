package main

import (
	"log/slog"
	"os"
	"time"
	"timeflow/internal/config"
	"timeflow/internal/data"
	"timeflow/internal/handlers"

	"gopkg.in/telebot.v4"
)

func main() {

	// Инициализация логгера
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Загрузка конфига
	cfg := config.Load()

	// Настройка бота
	pref := telebot.Settings{
		Token: cfg.Token,
		Poller: &telebot.LongPoller{
			Timeout: 5 * time.Second,
		},
	}

	// Создание бота
	bot, err := telebot.NewBot(pref)
	if err != nil {
		slog.Error("Error creating bot", "error", err)
		os.Exit(1)
	}

	// Регистрация всех обработчиков
	handlers.RegisterAll(bot)

	// Запуск бота
	slog.Info("Starting bot...")
	data.TAG()
	bot.Start()
}
