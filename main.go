package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var replyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Calatog"),
		tgbotapi.NewKeyboardButton("My profile"),
		tgbotapi.NewKeyboardButton("About this bot"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5532916297:AAGl7CvE3hs23M-iq6xoTtmpiKpK5tkASw8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Print(update)

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))
			bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Command() {
			case "start":
				msg.Text = update.Message.From.FirstName + ", welcome to LoremMarket!"
				msg.ReplyMarkup = replyKeyboard

			default:
				msg.Text = "Invalid response!"
			}

			bot.Send(msg)
		}
	}
}
