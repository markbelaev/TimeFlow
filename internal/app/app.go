package app

import (
	"log/slog"
	"os"
	"time"
	"timeflow/internal/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/telebot.v4"
)

// Настройки роутера
func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

// Настроки бота
func SetupBot(cfg *config.Config) (*telebot.Bot, error) {
	pref := telebot.Settings{
		Token: cfg.Token,
		Poller: &telebot.LongPoller{
			Timeout: 5 * time.Second,
		},
	}

	return telebot.NewBot(pref)
}

// Настройка логера
func SetupLooger() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Any().(time.Time)

				return slog.String("time", t.Format("15:04:05"))
			}

			return a
		},
	})))
}
