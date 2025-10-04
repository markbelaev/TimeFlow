package commands

import (
	"log/slog"
	"timeflow/internal/models"

	"gopkg.in/telebot.v4"
)

// Функция обработки команды /start
func HandleStart(c telebot.Context) error {

	// Получение информации об пользователе
	user := models.GetUserInfo(c.Sender())

	// Логирование комманды
	slog.Info("Обработка /start", "fist_name", user.FirstName, "last_name", user.LastName)

	// Текст для пользователя
	msg := "Привет, " + user.FirstName + " " + user.LastName + "!" + "\n\n" +
		"Это бот для твоего TimeFlow"

	// Возвращаем ответ
	return c.Send(msg)
}

// Обработчик команды /about
func HandleAbout(c telebot.Context) error {
	user := models.GetUserInfo(c.Sender())

	slog.Info("Обработка /about", "fist_name", user.FirstName, "last_name", user.LastName)

	msg := "Я разработчик"

	return c.Send(msg)
}
