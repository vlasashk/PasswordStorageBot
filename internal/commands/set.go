package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/configs"
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
	if CheckServiceLimit(user) {
		chatID := update.Message.Chat.ID
		msgID := update.Message.MessageID
		msg := update.Message.Text
		if users.UserData[userID].Input.Login {
			if len(msg) > 50 {
				client.Send(update.Message.Chat.ID, configs.LenErrMsg)
				delete(user.ServiceName, user.CurrServ)
				TerminateCommand(users, userID)
			} else {
				client.Send(update.Message.Chat.ID, configs.PassMsg+user.CurrServ)
				LoginFiller(users, msg, userID)
				go DeleteHandler(client.Bot, chatID, msgID)
				log.Println("Login initialized: ", users.UserData[userID])
			}
		} else if users.UserData[userID].Input.Pass {
			if len(msg) > 50 {
				client.Send(update.Message.Chat.ID, configs.LenErrMsg)
				delete(user.ServiceName, user.CurrServ)
				TerminateCommand(users, userID)
			} else {
				client.Send(update.Message.Chat.ID, configs.SuccessMsg+user.CurrServ)
				PasswordFiller(users, msg, userID)
				go DeleteHandler(client.Bot, chatID, msgID)
				log.Println("Password initialized: ", users.UserData[userID])
			}
		} else if _, ok := user.ServiceName[msg]; !ok {
			if len(msg) > 50 {
				client.Send(update.Message.Chat.ID, configs.LenErrMsg)
				TerminateCommand(users, userID)
			} else {
				client.Send(update.Message.Chat.ID, configs.LoginMsg+msg)
				storage.InitService(users, userID, msg)
				ServiceFiller(users, msg, userID)
				log.Println("Service initialized: ", users.UserData[userID])
			}
		} else {
			client.Send(update.Message.Chat.ID, msg+configs.ExistMsg)
			TerminateCommand(users, userID)
		}
	} else {
		client.Send(update.Message.Chat.ID, configs.LimitMsg)
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

func CheckServiceLimit(user storage.User) (res bool) {
	if len(user.ServiceName) < 20 {
		res = true
	}
	return
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
	service := user.ServiceName[user.CurrServ]
	service.Login = msg
	user.ServiceName[user.CurrServ] = service
	users.UserData[userID] = user
}

func PasswordFiller(users *storage.UserServices, msg string, userID int64) {
	user := users.UserData[userID]
	user.Input.Pass = false
	user.Input.Cmd = 0
	service := user.ServiceName[user.CurrServ]
	service.Password = msg
	user.ServiceName[user.CurrServ] = service
	users.UserData[userID] = user
}