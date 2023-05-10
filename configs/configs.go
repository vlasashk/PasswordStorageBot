package configs

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

var BotToken string = os.Getenv("TELEGRAM_BOT_TOKEN")
var SecretKey = []byte("UltraMeGaGigA_S3cr3tK3Y_No1_kNow")

type MessageConfig struct {
	tgbotapi.BaseChat
	Text                  string
	ParseMode             string
	DisableWebPagePreview bool
}
