package test

import (
	"fmt"
	"github.com/nntaoli-project/goex/v2"
	"github.com/nntaoli-project/goex/v2/model"
	"log"
	"testing"
)

func Test_Depth(t *testing.T) {
	client := goex.OKx.Spot

	_, _, err := client.GetExchangeInfo() //建议调用
	if err != nil {
		panic(err)
	}

	btcUSDTCurrencyPair, err := client.NewCurrencyPair(model.BTC, model.USDT)
	if err != nil {
		panic(err)
	}
	depth, _, err := client.GetDepth(btcUSDTCurrencyPair, 1)
	if err != nil {
		log.Fatalf("获取深度数据失败: %v", err)
	}
	fmt.Println("买单:", depth.Bids)
	fmt.Println("卖单:", depth.Asks)
}
