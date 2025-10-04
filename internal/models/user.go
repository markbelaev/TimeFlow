package models

import "gopkg.in/telebot.v4"

type TgUser struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
}

func GetUserInfo(user *telebot.User) TgUser {
	return TgUser{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}
}
