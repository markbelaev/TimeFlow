package commands

import (
	"log/slog"

	"gopkg.in/telebot.v4"
)

// Функция обработки комманды /start
func HandleStart(c telebot.Context) error {
	slog.Info("successful command processing /start")
	return c.Send("HELLO!")
}
