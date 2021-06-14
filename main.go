package main

import (
	"log"
	"reflect"
	"strings"

	"github.com/tkanos/gonfig"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Config with telegram token
type Config struct {
	TOKEN string
}

func main() {
	configuration := Config{}
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(configuration.TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			args := strings.Split(update.Message.Text, " ")
			switch args[0] {
			case "/start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hi bro")
				bot.Send(msg)

			case "/phone":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Started")
				bot.Send(msg)
				res := bomber(args[1])
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
				bot.Send(msg)

			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Use /phone to start bomber")
				bot.Send(msg)
			}

		}
	}
}
