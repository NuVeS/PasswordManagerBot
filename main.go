package main

import (
	"log"
	"strconv"

	"github.com/NuVeS/PasswordMangerBot/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5046824410:AAGR6jhP2X5v0kWDLg_cHIc3wRytHXHLUoQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for {
		update := <-updates

		if update.Message == nil {
			continue
		}

		chatId := strconv.FormatInt(update.Message.Chat.ID, 10)

		response := handler.HandleMessage(update.Message.Text, chatId)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		bot.Send(msg)
	}
}
