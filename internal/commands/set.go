package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/configs"
	"github.com/vlasashk/PasswordStorageBot/internal/security"
	"github.com/vlasashk/PasswordStorageBot/internal/storage"
	"log"
)

func (client *ClientConfig) SetInformer(update *tgbotapi.Update, users *storage.UserServices) {
	client.Send(update.Message.Chat.ID, configs.SetMsg)
	user := users.UserData[update.Message.From.ID]
	user.Input.Cmd = 1
	users.UserData[update.Message.From.ID] = user
}

func (client *ClientConfig) SetHandler(update *tgbotapi.Update, users *storage.UserServices) {
	userID := update.Message.From.ID
	user := users.UserData[userID]
	chatID := update.Message.Chat.ID
	msgID := update.Message.MessageID
	msg := update.Message.Text
	if limit := storage.CheckServiceLimit(client.DB, userID); limit && len(msg) <= 50 {
		if users.UserData[userID].Input.Login {
			client.Send(chatID, configs.PassMsg+user.CurrServ)
			LoginFiller(users, msg, userID)
			go security.DeleteHandler(client.Bot, chatID, msgID)
			log.Println("Login initialized: ", users.UserData[userID])
		} else if users.UserData[userID].Input.Pass {
			pass, err := security.EncryptPass(msg)
			if err != nil {
				client.Send(chatID, configs.EncrFailMsg)
			} else {
				client.Send(chatID, configs.SuccessMsg+user.CurrServ)
				PasswordFiller(users, pass, userID)
				storage.AddCredentials(client.DB, users.UserData[userID])
				go security.DeleteHandler(client.Bot, chatID, msgID)
				log.Println("Password initialized: ", users.UserData[userID])
			}
		} else if !storage.ServiceExist(client.DB, userID, msg) {
			client.Send(chatID, configs.LoginMsg+msg)
			ServiceFiller(users, msg, userID)
			log.Println("Service initialized: ", users.UserData[userID])
		} else {
			client.Send(chatID, msg+configs.ExistMsg)
			TerminateCommand(users, userID)
		}
	} else {
		if !limit {
			client.Send(chatID, configs.LimitMsg)
		} else {
			client.Send(chatID, configs.LenErrMsg)
			go security.DeleteHandler(client.Bot, chatID, msgID)
		}
		TerminateCommand(users, userID)
	}
}

func TerminateCommand(users *storage.UserServices, userID int64) {
	user := users.UserData[userID]
	user.Input.Login = false
	user.Input.Pass = false
	user.Input.Cmd = 0
	users.UserData[userID] = user
}

func ServiceFiller(users *storage.UserServices, msg string, userID int64) {
	user := users.UserData[userID]
	user.Input.Login = true
	user.CurrServ = msg
	users.UserData[userID] = user
}

func LoginFiller(users *storage.UserServices, msg string, userID int64) {
	user := users.UserData[userID]
	user.Input.Login = false
	user.Input.Pass = true
	user.ServiceName.Login = msg
	users.UserData[userID] = user
}

func PasswordFiller(users *storage.UserServices, msg string, userID int64) {
	user := users.UserData[userID]
	user.Input.Pass = false
	user.Input.Cmd = 0
	user.ServiceName.Password = msg
	users.UserData[userID] = user
}
