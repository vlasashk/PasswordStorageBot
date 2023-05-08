package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func DeleteHandler(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
	time.Sleep(1 * time.Minute)

	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := bot.Request(deleteMsg)
	if err != nil {
		log.Println("Failed to delete message:", err)
	}
}

func DeleteInlineKeyboard(bot *tgbotapi.BotAPI, chatID int64, messageID int, min int64) {
	time.Sleep(time.Duration(min) * time.Minute)
	emptyKeyboard := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{},
	}
	edit := tgbotapi.EditMessageReplyMarkupConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatID,
			MessageID:   messageID,
			ReplyMarkup: &emptyKeyboard,
		},
	}

	if _, err := bot.Send(edit); err != nil {
		log.Println("Failed to remove inline keyboard:", err)
	}
}
