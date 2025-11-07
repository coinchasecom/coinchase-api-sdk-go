package main

import (
	"fmt"

	"github.com/coinchasecom/coinchase-api-sdk-go/models"
	"github.com/coinchasecom/coinchase-api-sdk-go/services"
)

func main()  {
	// Buy
	payload := models.PlaceRequestParams{
		Volume: "10000",
		Price: "0.000057",
		Pair: "CCH/USDT",
	}
	buy := services.OrderBuy(payload)
	fmt.Println(buy)

	// Sell
	//payload = models.PlaceRequestParams{
	//	Volume: "10000",
	//	Price: "0.000057",
	//	Pair: "CCH/USDT",
	//}
	//sell := services.OrderSell(payload)
	//fmt.Println(sell)

	// Cancel One
	//id := "1205912407608659968"
	//kind := "1"
	//cancelOne := services.OrderCancelOne(id, kind)
	//fmt.Println(cancelOne)

	// Cancel by Pair
	//pair := "CCH/USDT"
	//cancelByPair := services.OrderCancelByPair(pair)
	//fmt.Println(cancelByPair)
}
