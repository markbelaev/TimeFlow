package commands

import (
	"log/slog"
	"timeflow/internal/data"

	"gopkg.in/telebot.v4"
)

// Функция обработки команды /start
func HandleStart(c telebot.Context) error {
	user := c.Sender()
	slog.Info("Обработка /start", "userID", user.ID)

	return c.Send(data.StartMessage)
}

// Обработчик команды /about
func HandleAbout(c telebot.Context) error {
	user := c.Sender()
	slog.Info("Обработка /about", "userID", user.ID)

	return c.Send(data.AboutMessage)
}
