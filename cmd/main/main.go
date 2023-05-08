package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/configs"
	"github.com/vlasashk/PasswordStorageBot/internal"
	"github.com/vlasashk/PasswordStorageBot/internal/commands"
	"github.com/vlasashk/PasswordStorageBot/internal/storage"
)

func main() {
	tgMenu := internal.InitKeyboard()
	var msg tgbotapi.MessageConfig
	usersData := storage.InitUsersStorage()

	client := commands.ClientConfig{
		Bot:  commands.InitBot(configs.BotToken),
		Menu: &tgMenu,
		Msg:  &msg}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := client.Bot.GetUpdatesChan(u)

	for update := range updates {
		var userID int64
		if update.Message != nil {
			userID = update.Message.From.ID
		} else if update.CallbackQuery != nil {
			userID = update.CallbackQuery.From.ID
		}
		storage.InitUser(usersData, userID)
		client.UpdateHandler(&update, usersData)
	}
}
