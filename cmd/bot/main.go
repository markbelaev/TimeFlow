package main

import (
	"log/slog"
	"os"
	"time"
	"timeflow/internal/config"

	"gopkg.in/telebot.v4"
)

func main() {

	// Инициализация логгера
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

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

	// Запуск бота
	slog.Info("Starting bot..")
	bot.Start()
}
