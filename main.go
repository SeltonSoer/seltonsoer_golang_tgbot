package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"seltonsoer_golang_tgbot/dbConnection"
	"seltonsoer_golang_tgbot/utils"
)

var devKey = "6329808132:AAH7wM9EIST0kKUU5Eo-DaHuvjZJfU9EHoU"
var prodKey = "6205520861:AAHQoorzrntpM2k1ExF3bHf8KFvAsLnQ8v0"

func main() {
	connectToTg()
}

func connectToTg() {
	bot, err := tgbotapi.NewBotAPI(prodKey)

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

		tgUser := update.Message.From
		bibaSize := getBibaSize()
		user := utils.User{
			UserName: tgUser.UserName,
			IdTgUser: tgUser.ID,
			BibaSize: bibaSize,
		}

		if update.Message.Text == "/biba_size" {
			userFromDb, errorSql := dbConnection.GetRecord(user)
			if errorSql == nil {
				msg := setFormMsgTg(update, userFromDb.BibaSize)
				bot.Send(msg)
			} else {
				log.Panic(errorSql)
			}
		}
		if update.Message.Text == "/refresh_biba" {
			_, errorSql := dbConnection.InsertConflictRecord(user)
			if errorSql == nil {
				msg := setFormMsgTg(update, user.BibaSize)
				bot.Send(msg)
			} else {
				log.Panic(errorSql)
			}
		}
	}
}

func getBibaSize() int {
	var randomNumber int
	randomNumber = utils.GetRandomNumberFromRange(0, 30)
	return randomNumber
}

func setFormMsgTg(update tgbotapi.Update, bibaSize int) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig
	var additionalMessage string = ""

	if bibaSize <= 12 {
		additionalMessage = "Печально наверно, с таким то малышом"
	} else if bibaSize <= 29 {
		additionalMessage = "Это уже достойно"
	} else {
		additionalMessage = "Вообще без комментариев))"
	}
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Твоя биба %s см. "+additionalMessage, fmt.Sprint(bibaSize)))

	return msg
}
