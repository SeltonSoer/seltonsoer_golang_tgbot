package main

import (
	"seltonsoer_golang_tgbot/utils"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6205520861:AAHQoorzrntpM2k1ExF3bHf8KFvAsLnQ8v0")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 10

	updates, _ := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/biba_size" {
			bot.Send(getBibaSize(update))
		}
	}

}

func getBibaSize(update tgbotapi.Update) tgbotapi.MessageConfig {

	var randomNumber int
	var msg tgbotapi.MessageConfig

	randomNumber = utils.GetRandomNumberFromRange(0, 30)
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Твоя биба %s см", fmt.Sprint(randomNumber)))

	return msg
}
