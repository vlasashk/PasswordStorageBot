package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/configs"
	"github.com/vlasashk/PasswordStorageBot/internal/storage"
	"strings"
)

func (client *ClientConfig) DelInformer(update *tgbotapi.Update, users *storage.UserServices) {
	userID := update.Message.From.ID
	userServices := storage.GetServicesByUser(client.DB, userID)
	if len(userServices) > 0 {
		sentMsg := client.SendWithKeyboard(update.Message.Chat.ID, configs.DelMsg, "del", userServices)
		go DeleteInlineKeyboard(client.Bot, sentMsg.Chat.ID, sentMsg.MessageID, 2)
		user := users.UserData[userID]
		user.Input.Cmd = 3
		users.UserData[userID] = user
	} else {
		client.Send(update.Message.Chat.ID, configs.EmptyMsg)
	}
}

func (client *ClientConfig) DelHandler(update *tgbotapi.Update, users *storage.UserServices) {
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
		_ = storage.DeleteService(client.DB, userID, msg)
		replyString := msg + " credentials have been successfully deleted"
		client.Send(chatID, replyString)
	} else {
		client.Send(chatID, configs.MissingMsg+msg)
	}
	TerminateCommand(users, userID)
}
