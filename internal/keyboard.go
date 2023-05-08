package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vlasashk/PasswordStorageBot/internal/storage"
)

func InitKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Set Password"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Get Pass"),
			tgbotapi.NewKeyboardButton("Delete Pass"),
		),
	)
}

func InitServicesKeyboard(services map[string]storage.Service, cmd string) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	var currentRow []tgbotapi.InlineKeyboardButton
	keys := make([]string, 0, len(services))
	for key := range services {
		keys = append(keys, key)
	}

	for i, key := range keys {
		button := tgbotapi.NewInlineKeyboardButtonData(key, key+" "+cmd)
		currentRow = append(currentRow, button)

		if len(currentRow) == 2 || i == len(keys)-1 {
			rows = append(rows, currentRow)
			currentRow = []tgbotapi.InlineKeyboardButton{}
		}
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)

	return keyboard
}
