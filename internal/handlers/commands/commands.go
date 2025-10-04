package commands

import (
	"log/slog"
	"strconv"
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
	msg := "Приветик! " + user.FirstName

	// Возвращаем ответ
	return c.Send(msg)
}

// Обработчик команды /about
func HandleAbout(c telebot.Context) error {
	user := models.GetUserInfo(c.Sender())

	slog.Info("Обработка /about", "fist_name", user.FirstName, "last_name", user.LastName)

	msg := "Я разработчик, а твой ID: " + strconv.Itoa(user.ID)

	return c.Send(msg)
}
