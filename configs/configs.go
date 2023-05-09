package configs

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const BotToken string = "YOUR_TOKEN_HERE"

type MessageConfig struct {
	tgbotapi.BaseChat
	Text                  string
	ParseMode             string
	DisableWebPagePreview bool
}
