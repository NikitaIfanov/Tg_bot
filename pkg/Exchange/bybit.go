package Exchange

func (p Pair) GetByBit() DataFloat {

	body := GetJson("https://api.bybit.com/v2/public/tickers?symbol=" + p.ByBit)

	targetsByBit := ByBitJson{}

	JsonUnmarshal(body, &targetsByBit)
	byBit := DataFloat{
		Exchange:  ByBit,
		Flag:      true,
		BuyPrice:  ParseFloat(targetsByBit.Data[0].BuyPrice),
		SalePrice: ParseFloat(targetsByBit.Data[0].SalePrice),
	}

	return byBit
}

type ByBitJson struct {
	Data []struct {
		Symbol    string `json:"symbol"`
		BuyPrice  string `json:"bid_price"`
		SalePrice string `json:"ask_price"`
	} `json:"result"`
}
