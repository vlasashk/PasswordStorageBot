package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/internal"
	"github.com/vlasashk/PasswordStorageBot/internal/storage"
	"gorm.io/gorm"
	"log"
)

type ClientConfig struct {
	DB   *gorm.DB
	Bot  *tgbotapi.BotAPI
	Menu *tgbotapi.ReplyKeyboardMarkup
	Msg  *tgbotapi.MessageConfig
}

func InitBot(token string) (bot *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return
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
