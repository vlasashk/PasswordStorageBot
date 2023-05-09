package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/configs"
	"github.com/vlasashk/PasswordStorageBot/internal/storage"
	"strings"
)

func (client *ClientConfig) GetInformer(update *tgbotapi.Update, users *storage.UserServices) {
	userID := update.Message.From.ID
	userServices := storage.GetServicesByUser(client.DB, userID)
	if len(userServices) > 0 {
		sentMsg := client.SendWithKeyboard(update.Message.Chat.ID, configs.GetMsg, "get", userServices)
		go DeleteInlineKeyboard(client.Bot, sentMsg.Chat.ID, sentMsg.MessageID, 2)
		user := users.UserData[userID]
		user.Input.Cmd = 2
		users.UserData[userID] = user
	} else {
		client.Send(update.Message.Chat.ID, configs.EmptyMsg)
	}
}

func (client *ClientConfig) GetHandler(update *tgbotapi.Update, users *storage.UserServices) {
	var userID int64
	var msg string
	var chatID int64
	if update.Message != nil {
		userID = update.Message.From.ID
		msg = update.Message.Text
		chatID = update.Message.Chat.ID
	} else {
		userID = update.CallbackQuery.From.ID
		msg = strings.Split(update.CallbackQuery.Data, " ")[0]
		chatID = update.CallbackQuery.Message.Chat.ID
	}
	ok := storage.ServiceExist(client.DB, userID, msg)
	if ok {
		service := storage.GetService(client.DB, userID, msg)
		replyString := msg + " credentials.\n" + "\nLogin: " + service.Login + "\nPassword: " + service.Password
		sentMsg := client.Send(chatID, replyString)
		go DeleteHandler(client.Bot, sentMsg.Chat.ID, sentMsg.MessageID)
	} else {
		client.Send(chatID, configs.MissingMsg+msg)
	}
	TerminateCommand(users, userID)
}
