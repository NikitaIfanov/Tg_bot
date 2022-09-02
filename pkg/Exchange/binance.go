package Exchange

func (p Pair) GetBinance() DataFloat {

	body := GetJson("https://api.binance.com/api/v3/ticker/bookTicker?symbol=" + p.Binance)

	targetsBinance := BinanceJson{}

	JsonUnmarshal(body, &targetsBinance)
	binance := DataFloat{
		Exchange:  Binance,
		Flag:      true,
		BuyPrice:  ParseFloat(targetsBinance.BuyPrice),
		SalePrice: ParseFloat(targetsBinance.SalePrice),
	}

	return binance

}

type BinanceJson struct {
	Symbol    string `json:"symbol"`
	BuyPrice  string `json:"bidPrice"`
	SalePrice string `json:"askPrice"`
}
