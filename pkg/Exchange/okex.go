package Exchange

func (p Pair) GetOkex() DataFloat {
	body := GetJson("https://www.okx.com/api/v5/market/ticker?instId=" + p.Okex + "-SWAP")

	targetsOkex := OkexJson{}

	JsonUnmarshal(body, &targetsOkex)
	okex := DataFloat{
		Exchange:  Okex,
		Flag:      true,
		BuyPrice:  ParseFloat(targetsOkex.Data[0].BuyPrice),
		SalePrice: ParseFloat(targetsOkex.Data[0].SalePrice),
	}

	return okex
}

type OkexJson struct {
	Data []struct {
		Symbol    string `json:"instId"`
		BuyPrice  string `json:"bidPx"`
		SalePrice string `json:"askPx"`
	}
}
