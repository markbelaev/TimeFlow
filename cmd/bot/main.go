package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
	"timeflow/internal/config"
	"timeflow/internal/data"
	"timeflow/internal/handlers"

	"github.com/gin-gonic/gin"
	"gopkg.in/telebot.v4"
)

func main() {

	// Показ тега
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

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	go func() {
		slog.Info("Starting API...")
		router.Run(":8080")
	}()

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
