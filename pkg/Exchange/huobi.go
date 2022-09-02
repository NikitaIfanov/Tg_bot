package Exchange

func (p Pair) GetHuobi() DataFloat {
	body := GetJson("https://api.huobi.pro/market/tickers")

	targetsHuobi := HuobiJson{}

	JsonUnmarshal(body, &targetsHuobi)
	huobi := DataFloat{
		Exchange: Huobi,
		Flag:     true,
	}
	for _, t := range targetsHuobi.Data {
		if t.Symbol == p.Huobi {
			huobi.BuyPrice = t.BuyPrice
			huobi.SalePrice = t.SalePrice
			break
		}
	}
	return huobi
}

type HuobiJson struct {
	Data []struct {
		Symbol    string  `json:"symbol"`
		BuyPrice  float64 `json:"bid"`
		SalePrice float64 `json:"ask"`
	}
}
