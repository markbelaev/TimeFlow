package handlers

import (
	"timeflow/internal/handlers/commands"

	"gopkg.in/telebot.v4"
)

// Функция для регистарции всех обработчиков
func RegisterAll(b *telebot.Bot) {
	registerCommands(b)
}

// Функция регистрации всех комманд
func registerCommands(b *telebot.Bot) {
	b.Handle("/start", commands.HandleStart)
	b.Handle("/about", commands.HandleAbout)
}
