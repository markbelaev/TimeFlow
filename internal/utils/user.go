package utils

import (
	"timeflow/internal/models"

	"gopkg.in/telebot.v4"
)

// Функция получения информации о пользователе
func GetUserInfo(user *telebot.User) models.TgUser {
	return models.TgUser{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}
}
