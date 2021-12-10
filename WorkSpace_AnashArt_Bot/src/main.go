/*
 ------------------- ЗАДАЧИ -------------------
1) Добавить Callback кнопки в START MESSAGE
2) Изменять текст при нажатии на Callback
3) Записывать логи в файл
 ----------------------------------------------
*/

package main

// ------------------- IMPORTS -------------------
import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ------------------- CONSTS -------------------
const botAPI = "5015857552:AAGHOqwNAeeJ4Su0Rnlu9UOAz6MaO3IDpng"

// ------------------- KEYBOARDS-------------------
var startMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Перейти на AnashArt.ru", "https://AnashArt.ru"),
		tgbotapi.NewInlineKeyboardButtonData("Показать Коллекции", "0"),
	),
)

// ------------------- FUNCS -------------------
func main() {
	// --------- PROJECT INIT --------
	logFile, err := os.Create("logs.txt")
	if err != nil {
		log.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer logFile.Close()

	// --------- INIT BOT ---------
	bot, err := tgbotapi.NewBotAPI(botAPI)
	if err != nil {
		logFile.WriteString("ERROR: " + err.Error() + "\n")
		log.Panic(err)
	}
	bot.Debug = true
	logFile.WriteString("Authorized on account " + bot.Self.UserName + "\n")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// --------- CHECK NEW MESSAGE LOOP ---------
	for update := range updates {

		// --------- STANDART MESSAGE ---------
		if update.Message != nil {

			// --------- COMMAND MESSAGE ---------
			if update.Message.IsCommand() {

				// --------- START MESSAGE ---------
				if update.Message.Command() == "start" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я бот AnashArt.\n\nТы можешь перейти на наш сайт или попросить меня показать все наши коллекции")
					msg.ReplyMarkup = startMenu
					if _, err = bot.Send(msg); err != nil {
						logFile.WriteString("ERROR: " + err.Error() + "\n")
						panic(err)
					}

					logFile.WriteString("New Message: " + update.Message.Text + "\n")
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я даже не знаю как на это ответить(")
				bot.Send(msg)

				logFile.WriteString("New Message: " + update.Message.Text + "\n")
			}
		}

		// --------- CALLBACK MESSAGE ---------
		if update.CallbackQuery != nil {

			switch update.CallbackQuery.Data {

			// --------- CALLBACK START MESSAGE ---------
			case "0":
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := bot.Request(callback); err != nil {
					logFile.WriteString("ERROR: " + err.Error() + "\n")
					panic(err)
				}

				// msgConf := tgbotapi.NewEditMessageText(chatid, messageid, "ПЫЩЩЩЩЩЩ")
				// update.EditedMessage.MigrateToChatID = msgConf.ChatID

			case "1":
				// CODE ...

				// ------------------- НАРАБОТКИ -------------------

				// callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				// if _, err := bot.Request(callback); err != nil {
				// 	panic(err)
				// }

				// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				// bot.Send(msg)
			}
		}
	}
}
