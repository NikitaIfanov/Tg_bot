package Exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	All     = "All"
	Binance = "Binance"
	Huobi   = "Huobi"
	Okex    = "Okex"
	ByBit   = "ByBit"
	Kraken  = "Kraken"
)

type DataFloat struct {
	Exchange  string
	Flag      bool
	BuyPrice  float64
	SalePrice float64
}
type SelectExchange struct {
	All         bool
	BinanceTrue bool
	HuobiTrue   bool
	OkexTrue    bool
	ByBitTrue   bool
	KrakenTrue  bool
}

func JsonUnmarshal(body []byte, targets interface{}) {
	err := json.Unmarshal(body, &targets)
	if err != nil {
		log.Fatal(err)
	}
}

func GetJson(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return body
}

func (p Pair) Make(s SelectExchange) []DataFloat {
	Data := make([]DataFloat, 0, 5)

	if s.All == true {
		Data = append(Data, p.GetBinance(), p.GetHuobi(), p.GetOkex(), p.GetByBit(), p.GetKraken())
		return Data
	}

	if s.BinanceTrue == true {
		Data = append(Data, p.GetBinance())
	}
	if s.HuobiTrue == true {
		Data = append(Data, p.GetHuobi())
	}
	if s.OkexTrue == true {
		Data = append(Data, p.GetOkex())
	}
	if s.ByBitTrue == true {
		Data = append(Data, p.GetByBit())
	}
	if s.KrakenTrue == true {
		Data = append(Data, p.GetKraken())
	}

	return Data
}

type Pair struct {
	Difference float64
	Binance    string
	Huobi      string
	Okex       string
	ByBit      string
	Kraken     string
}

func EnterPair(FirstExchange string, SecondExchange string, Difference string) Pair {

	return Pair{
		Difference: ParseFloat(Difference),
		Binance:    fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
		Huobi:      fmt.Sprintf("%s%s", strings.ToLower(FirstExchange), strings.ToLower(SecondExchange)),
		Okex:       fmt.Sprintf("%s-%s", FirstExchange, SecondExchange),
		ByBit:      fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
		Kraken:     fmt.Sprintf("%s%s", FirstExchange, SecondExchange),
	}

}
func ParseFloat(s string) float64 {
	data, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

type TrackingPair struct {
	Flag    bool
	ID      int
	MinBuy  DataFloat
	MaxSale DataFloat
}

func Tracking(Data []DataFloat, difference float64) TrackingPair {
	t := TrackingPair{
		Flag:    false,
		ID:      rand.Int(),
		MinBuy:  Data[0],
		MaxSale: Data[0],
	}

	for _, num := range Data {
		switch {
		case num.SalePrice > t.MaxSale.SalePrice:
			t.MaxSale = num

		case num.BuyPrice < t.MinBuy.BuyPrice:
			t.MinBuy = num

		}
	}
	if t.MaxSale.SalePrice-t.MinBuy.BuyPrice > difference {
		t.Flag = true
		return t
	}
	return t
}
