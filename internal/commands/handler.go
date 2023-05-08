package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/configs"
	"github.com/vlasashk/PasswordStorageBot/internal"
	"github.com/vlasashk/PasswordStorageBot/internal/storage"
	"log"
	"strings"
)

type ClientConfig struct {
	Bot  *tgbotapi.BotAPI
	Menu *tgbotapi.ReplyKeyboardMarkup
	Msg  *tgbotapi.MessageConfig
}

func (client *ClientConfig) UpdateHandler(update *tgbotapi.Update, users *storage.UserServices) {
	if update.Message != nil { // ignore non-Message updates
		userID := update.Message.From.ID
		if update.Message.IsCommand() || client.KeyboardVerify(update) {
			TerminateCommand(users, userID)
			client.CommandHandler(update, users)
		} else {
			switch users.UserData[userID].Input.Cmd {
			case 0:
				client.CommandHandler(update, users)
			case 1:
				client.SetHandler(update, users)
			case 2:
				client.GetHandler(update, users)
			case 3:
				client.DelHandler(update, users)
			}
		}
	} else if update.CallbackQuery != nil {
		chatID := update.CallbackQuery.Message.Chat.ID
		msgId := update.CallbackQuery.Message.MessageID
		//userID := update.CallbackQuery.From.ID
		DeleteInlineKeyboard(client.Bot, chatID, msgId, 0)

		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if _, err := client.Bot.Request(callback); err != nil {
			panic(err)
		}
		cmd := strings.Split(update.CallbackQuery.Data, " ")[1]
		switch cmd {
		case "get":
			client.GetHandler(update, users)
		case "del":
			client.DelHandler(update, users)
		}
	}
}

func (client *ClientConfig) CommandHandler(update *tgbotapi.Update, users *storage.UserServices) {
	cmd := update.Message.Command()
	chatID := update.Message.Chat.ID
	client.KeyboardHandler(update, &cmd)

	switch cmd {
	case "start":
		*client.Msg = tgbotapi.NewMessage(update.Message.Chat.ID, configs.StartMsg)
		client.Msg.ReplyMarkup = client.Menu
		_, err := client.Bot.Send(*client.Msg)
		if err != nil {
			log.Println(err)
		}
		return
	case "set":
		client.SetInformer(update, users)
	case "get":
		client.GetInformer(update, users)
	case "del":
		client.DelInformer(update, users)
	case "help":
		client.Send(chatID, configs.HelpMsg)
	default:
		client.Send(chatID, configs.DefaultMsg)
	}
}

func (client *ClientConfig) KeyboardHandler(update *tgbotapi.Update, cmd *string) bool {
	res := true
	switch update.Message.Text {
	case client.Menu.Keyboard[0][0].Text:
		*cmd = "set"
	case client.Menu.Keyboard[1][0].Text:
		*cmd = "get"
	case client.Menu.Keyboard[1][1].Text:
		*cmd = "del"
	default:
		res = false
	}
	return res
}

func (client *ClientConfig) KeyboardVerify(update *tgbotapi.Update) bool {
	res := false
	switch update.Message.Text {
	case client.Menu.Keyboard[0][0].Text:
		res = true
	case client.Menu.Keyboard[1][0].Text:
		res = true
	case client.Menu.Keyboard[1][1].Text:
		res = true
	}
	return res
}

func (client *ClientConfig) Send(id int64, msg string) (sentMsg tgbotapi.Message) {
	*client.Msg = tgbotapi.NewMessage(id, msg)
	sentMsg, err := client.Bot.Send(*client.Msg)
	if err != nil {
		log.Println(err)
	}
	return
}

func (client *ClientConfig) SendWithKeyboard(id int64, msg string, cmd string, services map[string]storage.Service) (sentMsg tgbotapi.Message) {
	temp := tgbotapi.NewMessage(id, msg)
	temp.ReplyMarkup = internal.InitServicesKeyboard(services, cmd)
	*client.Msg = temp
	sentMsg, err := client.Bot.Send(*client.Msg)
	if err != nil {
		log.Println(err)
	}
	return
}