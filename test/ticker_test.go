package test

import (
	"fmt"
	"github.com/nntaoli-project/goex/v2/model"
	"testing"

	goex "github.com/nntaoli-project/goex/v2"
)

func Test_Ticker(t *testing.T) {
	client := goex.OKx.Spot

	_, _, err := client.GetExchangeInfo() //建议调用
	if err != nil {
		panic(err)
	}

	btcUSDTCurrencyPair, err := client.NewCurrencyPair(model.BTC, model.USDT)
	if err != nil {
		panic(err)
	}

	ticker, _, err := client.GetTicker(btcUSDTCurrencyPair)
	if err != nil {
		return
	}
	fmt.Printf("Last : %+v \n", ticker.Last)
	fmt.Printf("Buy : %+v \n", ticker.Buy)
	fmt.Printf("Sell : %+v \n", ticker.Sell)
	fmt.Printf("High : %+v \n", ticker.High)
	fmt.Printf("Low : %+v \n", ticker.Low)
}
