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

	// ТЭГ в консоли с именем бота
	data.Tag()

	// Инициализация логера
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Any().(time.Time)

				return slog.String("time", t.Format("15:04:05"))
			}

			return a
		},
	})))

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
	bot.Start()
}
