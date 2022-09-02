package tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"main.go/pkg/Exchange"
)

func Button(sTruth *Exchange.SelectExchange, flag string) string {
	switch flag {
	case Exchange.All:
		if sTruth.All == true {
			return Exchange.All + "✅"
		} else {
			return Exchange.All
		}
	case Exchange.Binance:
		if sTruth.BinanceTrue == true {
			return Exchange.Binance + "✅"
		} else {
			return Exchange.Binance
		}
	case Exchange.Huobi:
		if sTruth.HuobiTrue == true {
			return Exchange.Huobi + "✅"
		} else {
			return Exchange.Huobi
		}
	case Exchange.Okex:
		if sTruth.OkexTrue == true {
			return Exchange.Okex + "✅"
		} else {
			return Exchange.Okex
		}
	case Exchange.ByBit:
		if sTruth.ByBitTrue == true {
			return Exchange.ByBit + "✅"
		} else {
			return Exchange.ByBit
		}
	case Exchange.Kraken:
		if sTruth.KrakenTrue == true {
			return Exchange.Kraken + "✅"
		} else {
			return Exchange.Kraken
		}

	}
	return "Error"
}

func IncorrectCommandMsg(bot *tgbotapi.BotAPI, ChatID int64) {
	msg := tgbotapi.NewMessage(ChatID, "Incorrect command")
	if _, err := bot.Send(msg); err != nil {
		log.Print(err)
	}
}

func SendMsg(bot *tgbotapi.BotAPI, ChatID int64, message string) {
	msg := tgbotapi.NewMessage(ChatID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Print(err)
	}
}

func SendMsgWithKeyboard(message string, bot *tgbotapi.BotAPI, ChatID int64, keyboard tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(ChatID, message)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		log.Print(err)
	}
}
func ButtonChange(bot *tgbotapi.BotAPI, chatID int64, messageID int, keyboard tgbotapi.InlineKeyboardMarkup) {
	conf := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, keyboard)
	if _, err := bot.Send(conf); err != nil {
		log.Print(err)
	}
}

func SwitchAll(s *Exchange.SelectExchange) bool {
	if s.All == false {
		s.All = true
		s.BinanceTrue = true
		s.HuobiTrue = true
		s.OkexTrue = true
		s.ByBitTrue = true
		s.KrakenTrue = true
		return true
	} else {
		s.All = false
		s.BinanceTrue = false
		s.HuobiTrue = false
		s.OkexTrue = false
		s.ByBitTrue = false
		s.KrakenTrue = false
		return false
	}
}
func DataFloatToMsg(bot *tgbotapi.BotAPI, ChatID int64, data Exchange.DataFloat) {
	msg := fmt.Sprintf("%s:\nSale price: %.2f$\nBuy price: %.2f$", data.Exchange, data.SalePrice, data.BuyPrice)
	SendMsg(bot, ChatID, msg)
}
