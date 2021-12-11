/*
 ------------------- ЗАДАЧИ -------------------
1) Добавлять в структуру данные при CallBack
2) После заполнения всех полей отправлять структуру администратору
3) Создать Администрарора
4) Добавить команды через отца
5) Добавить оплату по карте
 ----------------------------------------------
*/

package main

// ------------------- IMPORTS -------------------
import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ------------------- CONSTS -------------------
const botAPI = "5015857552:AAGHOqwNAeeJ4Su0Rnlu9UOAz6MaO3IDpng"

const initMsg = "Привет! Я бот AnashArt и вот что я умею:\n\n  /show - покажу все наши коллекции\n  /price - покажу прайслист\n  /order - оформлю заказ\n  /operator - позову администратора\n\nНаш официальный сайт: https://AnashArt.ru\nНаш официальный Instagram: @Anash.Art"
const octopusPATH = "img/octopus.jpg"
const shrimpPATH = "img/shrimp.jpg"

// ------------------- KEYBOARDS-------------------
var TypeOrder = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Заказать в Telegram", "Telegram"),
		tgbotapi.NewInlineKeyboardButtonURL("Заказать в Insagram", "https://www.instagram.com/anash.art/"),
	),
)

var LotOrder = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("МОКРИЙ", "МОКРИЙ"),
		tgbotapi.NewInlineKeyboardButtonData("КРЕВЭД", "КРЕВЭД"),
	),
)

var SizeOrder = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("S", "S"),
		tgbotapi.NewInlineKeyboardButtonData("M", "M"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("L", "L"),
		tgbotapi.NewInlineKeyboardButtonData("XL", "XL"),
	),
)

// ------------------- STRUCTS ------------------
type OrderInfo struct {
	UserName string
	Email    string
	Print    string
	Size     string
	Addres   string
}

var OrderInfoMap map[int64]*OrderInfo

// ------------------- FUNCS -------------------
func init() {
	OrderInfoMap = make(map[int64]*OrderInfo)
}

func main() {

	// --------- INIT BOT ---------
	bot, err := tgbotapi.NewBotAPI(botAPI)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// --------- CHECK NEW MESSAGE LOOP ---------
	for update := range updates {

		// --------- STANDART MESSAGE ---------
		if update.Message != nil {

			// --------- COMMAND MESSAGE ---------
			if update.Message.IsCommand() {

				// --------- COMMAND START---------
				if update.Message.Command() == "start" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, initMsg)
					if _, err = bot.Send(msg); err != nil {
						panic(err)
					}

					// --------- COMMAND SHOW---------
				} else if update.Message.Command() == "show" {
					photoBytes, err := ioutil.ReadFile(octopusPATH)
					if err != nil {
						panic(err)
					}
					photoFileBytes := tgbotapi.FileBytes{
						Name:  "picture",
						Bytes: photoBytes,
					}
					bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, photoFileBytes))
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Принт МОКРИЙ"))

					photoBytes, err = ioutil.ReadFile(shrimpPATH)
					if err != nil {
						panic(err)
					}
					photoFileBytes = tgbotapi.FileBytes{
						Name:  "picture",
						Bytes: photoBytes,
					}
					bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, photoFileBytes))
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Принт КРЭВЕД"))

					// --------- COMMAND OPERATOR---------
				} else if update.Message.Command() == "operator" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В скором времени с вами свяжется наш администратор ...")
					if _, err = bot.Send(msg); err != nil {
						panic(err)
					}

					bot.Send(tgbotapi.NewMessage(853634511, "PROBLEM @"+update.Message.From.UserName))

					// --------- COMMAND PRICE---------
				} else if update.Message.Command() == "price" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Price List:\n\n  Футболка 'МОКРИЙ' - 3590\n  Футболка 'КРЕВЭД' - 3590")
					if _, err = bot.Send(msg); err != nil {
						panic(err)
					}

					// --------- COMMAND ORDER---------
				} else if update.Message.Command() == "order" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы можете заказать что либо прямо здесь или обратиться к администратору в нашем Instagram (@Anash.Art)")
					msg.ReplyMarkup = TypeOrder
					if _, err = bot.Send(msg); err != nil {
						panic(err)
					}

				}
				// --------- UNKNOWN MESSAGE ---------
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я даже не знаю как на это ответить(")
				bot.Send(msg)
			}
		}

		// --------- CALLBACK MESSAGE ---------
		if update.CallbackQuery != nil {

			switch update.CallbackQuery.Data {

			// --------- CALLBACK TYPE ORDER ---------
			case "Telegram":
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выбирай какой принт ты хочешь")
				msg.ReplyMarkup = LotOrder

				bot.Send(msg)

			// --------- CALLBACK LOT ORDER---------
			case "МОКРИЙ":
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Теперь нужно выбрать свой размер")
				msg.ReplyMarkup = SizeOrder

				bot.Send(msg)

			// --------- CALLBACK LOT ORDER---------
			case "КРЕВЭД":
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Теперь нужно выбрать свой размер")
				msg.ReplyMarkup = SizeOrder

				bot.Send(msg)

			// --------- CALLBACK SIZE ORDER---------
			case "S":

			// --------- CALLBACK SIZE ORDER---------
			case "M":

			// --------- CALLBACK SIZE ORDER---------
			case "L":

			// --------- CALLBACK SIZE ORDER---------
			case "XL":

			}
		}
	}
}
