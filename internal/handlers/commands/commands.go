package commands

import (
	"log/slog"
	"timeflow/internal/data"

	"gopkg.in/telebot.v4"
)

// Функция обработки комманды /start
func HandleStart(c telebot.Context) error {
	slog.Info("successful command processing /start")
	return c.Send(data.StartMessage)
}

// Обработчик комманды /about
func HandleAbout(c telebot.Context) error {
	slog.Info("successful command processing /about")
	return c.Send(data.AboutMessage)
}
