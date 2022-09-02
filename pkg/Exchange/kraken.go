package Exchange

func (p Pair) GetKraken() DataFloat {
	altName := make(map[string]string, 100)
	altName["ETHUSDT"] = "ETHUSDT"
	body := GetJson("https://api.kraken.com/0/public/Ticker?pair=" + p.Kraken)

	targetsKraken := KrakenJson{}

	JsonUnmarshal(body, &targetsKraken)

	kraken := DataFloat{
		Exchange:  Kraken,
		Flag:      true,
		BuyPrice:  ParseFloat(targetsKraken.Result[altName[p.Kraken]].BuyPrice[0]),
		SalePrice: ParseFloat(targetsKraken.Result[altName[p.Kraken]].SalePrice[0]),
	}
	return kraken
}

type KrakenJson struct {
	Result map[string]struct {
		BuyPrice  []string `json:"b"`
		SalePrice []string `json:"a"`
	} `json:"result"`
}
