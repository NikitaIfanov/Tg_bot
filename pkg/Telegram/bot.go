package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"main.go/configs"
	"main.go/pkg/Db"
	"main.go/pkg/Exchange"
	"strings"
	"time"
)

var s = Exchange.SelectExchange{
	BinanceTrue: false,
	HuobiTrue:   false,
	OkexTrue:    false,
	ByBitTrue:   false,
	KrakenTrue:  false,
}
var NewKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("New Pair", "newPair"),
		tgbotapi.NewInlineKeyboardButtonData("Show", "show"),
		tgbotapi.NewInlineKeyboardButtonData("Help", "help"),
	),
)
var ExchangeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Binance", "Binance"),
		tgbotapi.NewInlineKeyboardButtonData("Huobi", "Huobi"),
		tgbotapi.NewInlineKeyboardButtonData("Okex", "Okex"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ByBit", "ByBit"),
		tgbotapi.NewInlineKeyboardButtonData("Kraken(Not working)", "Kraken"),
		tgbotapi.NewInlineKeyboardButtonData("All", "All"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("See Price", "See"),
		tgbotapi.NewInlineKeyboardButtonData("Track Difference", "Track"),
	),
)

func Bot() {
	bot, err := tgbotapi.NewBotAPI(configs.MyBotToken) //Бот

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	var (
		pair         = Exchange.Pair{}
		trackingPair = Exchange.TrackingPair{}

		m    = make(map[string]bool)
		done = make(chan struct{})
	)

	for update := range updates {

		if update.Message != nil {
			switch {
			//Enter in bot
			case update.Message.Text == "/start":
				SendMsgWithKeyboard("Hello", bot, update.Message.Chat.ID, NewKeyboard)
			case update.Message.Text == "stop":
				done <- struct{}{}
			case len(update.Message.Text) >= 6:

				command := strings.Split(update.Message.Text, " ")
				switch {
				case len(command) == 2:
					pair = Exchange.EnterPair(command[0], command[1], "0")

					SendMsgWithKeyboard("Select Exchanges", bot, update.Message.Chat.ID, ExchangeKeyboard)
				case len(command) == 3:
					trackingPair.Pair = command[0] + " " + command[1]
					pair = Exchange.EnterPair(command[0], command[1], command[2])

					m[command[0]+command[1]+" "+command[2]] = true
					SendMsgWithKeyboard("Select Exchanges", bot, update.Message.Chat.ID, ExchangeKeyboard)
				default:

					SendMsg(bot, update.Message.Chat.ID, "Error, incorrect format, check the example")
					continue
				}

			default:
				IncorrectCommandMsg(bot, update.Message.Chat.ID)

			}

			//Exchange selection for the entered pair

		}

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "newPair":
				SendMsg(bot, update.CallbackQuery.Message.Chat.ID,
					"Enter a Cryptocurrency Pair Example:'BTC USDT' for see one time\n"+
						" or\n"+
						" Pair & Difference Example:'BTC USDT 10' for tracking")
			case "Binance":
				if s.BinanceTrue == false {
					s.BinanceTrue = true

					ExchangeKeyboard.InlineKeyboard[0][0].Text = Button(&s, Exchange.Binance)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Selected"+" "+Exchange.Binance)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)

				} else {
					s.All = false
					s.BinanceTrue = false
					ExchangeKeyboard.InlineKeyboard[0][0].Text = Button(&s, Exchange.Binance)
					ExchangeKeyboard.InlineKeyboard[1][2].Text = Button(&s, Exchange.All)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Canceled"+" "+Exchange.Binance)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
				}

			case "Huobi":
				if s.HuobiTrue == false {
					s.HuobiTrue = true
					ExchangeKeyboard.InlineKeyboard[0][1].Text = Button(&s, Exchange.Huobi)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Selected"+" "+Exchange.Huobi)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
				} else {
					s.All = false
					s.HuobiTrue = false
					ExchangeKeyboard.InlineKeyboard[0][1].Text = Button(&s, Exchange.Huobi)
					ExchangeKeyboard.InlineKeyboard[1][2].Text = Button(&s, Exchange.All)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Canceled"+" "+Exchange.Huobi)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)

				}

			case "Okex":
				if s.OkexTrue == false {
					s.OkexTrue = true
					ExchangeKeyboard.InlineKeyboard[0][2].Text = Button(&s, Exchange.Okex)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Selected"+" "+Exchange.Okex)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
				} else {
					s.All = false
					s.OkexTrue = false
					ExchangeKeyboard.InlineKeyboard[0][2].Text = Button(&s, Exchange.Okex)
					ExchangeKeyboard.InlineKeyboard[1][2].Text = Button(&s, Exchange.All)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Canceled"+" "+Exchange.Okex)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
				}

			case "ByBit":
				if s.ByBitTrue == false {
					s.ByBitTrue = true

					ExchangeKeyboard.InlineKeyboard[1][0].Text = Button(&s, Exchange.ByBit)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Selected"+" "+Exchange.ByBit)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
				} else {
					s.All = false
					s.ByBitTrue = false

					ExchangeKeyboard.InlineKeyboard[1][0].Text = Button(&s, Exchange.ByBit)
					ExchangeKeyboard.InlineKeyboard[1][2].Text = Button(&s, Exchange.All)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Canceled"+" "+Exchange.ByBit)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
				}
				//Kraken
			/*case "Kraken":
			if s.KrakenTrue == false {
				s.KrakenTrue = true
				ExchangeKeyboard.InlineKeyboard[1][1].Text = Button(&s, Exchange.Kraken)
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Selected"+" "+Exchange.Kraken)
				if _, err := bot.Request(callback); err != nil {
					log.Print(err)
				}
				ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
			} else {
				s.All = false
				s.KrakenTrue = false
				ExchangeKeyboard.InlineKeyboard[1][1].Text = Button(&s, Exchange.Kraken)
				ExchangeKeyboard.InlineKeyboard[1][2].Text = Button(&s, Exchange.All)
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Canceled"+" "+Exchange.Kraken)
				if _, err := bot.Request(callback); err != nil {
					log.Print(err)
				}
				ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)
			}*/

			case "All":
				if s.All == false {
					SwitchAll(&s)
					ExchangeKeyboard.InlineKeyboard[1][2].Text = Button(&s, Exchange.All)
					ExchangeKeyboard.InlineKeyboard[0][0].Text = Button(&s, Exchange.Binance)
					ExchangeKeyboard.InlineKeyboard[0][1].Text = Button(&s, Exchange.Huobi)
					ExchangeKeyboard.InlineKeyboard[0][2].Text = Button(&s, Exchange.Okex)
					ExchangeKeyboard.InlineKeyboard[1][0].Text = Button(&s, Exchange.ByBit)
					ExchangeKeyboard.InlineKeyboard[1][1].Text = Button(&s, Exchange.Kraken)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Selected"+" "+Exchange.All)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}
					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)

				} else {
					SwitchAll(&s)
					ExchangeKeyboard.InlineKeyboard[1][2].Text = Button(&s, Exchange.All)
					ExchangeKeyboard.InlineKeyboard[0][0].Text = Button(&s, Exchange.Binance)
					ExchangeKeyboard.InlineKeyboard[0][1].Text = Button(&s, Exchange.Huobi)
					ExchangeKeyboard.InlineKeyboard[0][2].Text = Button(&s, Exchange.Okex)
					ExchangeKeyboard.InlineKeyboard[1][0].Text = Button(&s, Exchange.ByBit)
					ExchangeKeyboard.InlineKeyboard[1][1].Text = Button(&s, Exchange.Kraken)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Canceled"+" "+Exchange.All)
					if _, err := bot.Request(callback); err != nil {
						log.Print(err)
					}

					ButtonChange(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, ExchangeKeyboard)

				}
			case "See":

				data := pair.Make(s)
				for _, d := range data {
					DataFloatToMsg(bot, update.CallbackQuery.Message.Chat.ID, d)
				}

			case "Track":

				go func(d chan struct{}, ChatID int64) {
					timer := time.NewTicker(1 * time.Second)
					trackingPair = Exchange.Tracking(pair.Make(s), pair.Difference)

					for {
						select {
						case <-d:
							timer.Stop()
							return
						case <-timer.C:

							if trackingPair.Flag == true {
								DataFloatToMsg(bot, ChatID, trackingPair.MinBuy)
								DataFloatToMsg(bot, ChatID, trackingPair.MaxSale)
							}

						}
					}

				}(done, update.CallbackQuery.Message.Chat.ID)

			case "show":

				SendMsg(bot, update.CallbackQuery.Message.Chat.ID, Db.Show(m))
				log.Print(m)

			case "help":
				SendMsg(bot, update.CallbackQuery.Message.Chat.ID, "Help")
			case "about":
				SendMsg(bot, update.CallbackQuery.Message.Chat.ID, "Test bot")
			default:
				IncorrectCommandMsg(bot, update.CallbackQuery.Message.Chat.ID)
			}

		}
	}
}
