package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const botAPI = "5046696065:AAGtZ2vc3__XVHTHQ9vkt2ytW1QApc94-c8"

// --------db--------
// type wallet map[string]float64

// var db = map[int64]wallet{}

// // --------json response--------
type binanceResponse struct {
	Price float64 `json:"price,string"`
	Code  int64   `json:"code"`
}

type NewToken struct {
	State int // 0 - Tiker; 1 - Value
	ID    int64
	Tiker string
	Value float64
}

var UsersWallet map[int64]*NewToken

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Add token"),
		tgbotapi.NewKeyboardButton("Delete token"),
		tgbotapi.NewKeyboardButton("Show tokens"),
		tgbotapi.NewKeyboardButton("Go back"),
	),
)

func init() {
	UsersWallet = make(map[int64]*NewToken)
}

func main() {
	var (
		bot *tgbotapi.BotAPI
		err error

		update        tgbotapi.Update
		updateChannel tgbotapi.UpdatesChannel
		UpdateConfig  tgbotapi.UpdateConfig
	)

	bot, err = tgbotapi.NewBotAPI(botAPI)
	if err != nil {
		panic("BOT INIT ERROR: " + err.Error())
	}

	UpdateConfig.Timeout = 60
	UpdateConfig.Limit = 1
	UpdateConfig.Offset = 0

	updateChannel = bot.GetUpdatesChan(UpdateConfig)

	for {
		update = <-updateChannel

		if update.Message.IsCommand() {

			if update.Message.Command() == "menu" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
				msg.ReplyMarkup = mainMenu
				bot.Send(msg)
			}

		} else {

			if update.Message.Text == mainMenu.Keyboard[0][0].Text {

				UsersWallet[update.Message.From.ID] = new(NewToken)
				UsersWallet[update.Message.From.ID].State = 0

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Token Ticker: ")
				bot.Send(msg)
			} else if update.Message.Text == mainMenu.Keyboard[0][1].Text {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Token Ticker: ")
				bot.Send(msg)
			} else if update.Message.Text == mainMenu.Keyboard[0][2].Text {
				// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter Token Ticker: ")
				// resp, err := getPrice(msg.Text)

				// if err != nil {
				// 	log.Panic("TOKEN NAME ERROR: " + err.Error())
				// }
				// bot.Send()
			} else if update.Message.Text == mainMenu.Keyboard[0][3].Text {
				// убрать меню
			}

		}
		fmt.Printf("From: %v; ChatID: %v; Messege: %v\n",
			update.Message.From.UserName,
			update.Message.Chat.ID,
			update.Message.Text)
	}
}

func getPrice(symbol string) (price float64, err error) {
	respone, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", symbol))
	if err != nil {
		return
	}

	defer respone.Body.Close()

	var jsonResponse binanceResponse
	err = json.NewDecoder(respone.Body).Decode(&jsonResponse)
	if err != nil {
		return
	}

	if jsonResponse.Code != 0 {
		err = errors.New("Неверный токен")
	}

	price = jsonResponse.Price

	return price, err
}
