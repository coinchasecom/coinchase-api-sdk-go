package services

import (
	"encoding/json"
	"net/url"

	"github.com/coinchasecom/coinchase-api-sdk-go/models"
	"github.com/coinchasecom/coinchase-api-sdk-go/utils"
)

// OrderBuy Place a buy order.
func OrderBuy(placeRequestParams models.PlaceRequestParams) models.PlaceReturn {
	placeReturn := models.PlaceReturn{}

	mapParams := make(map[string]string)
	mapParams["volume"] = placeRequestParams.Volume
	if 0 < len(placeRequestParams.Price) {
		mapParams["price"] = placeRequestParams.Price
	}
	mapParams["pair"] = placeRequestParams.Pair

	strRequest := "/x/v1/order/buy"
	jsonPlaceReturn := utils.ApiKeyPost(mapParams, strRequest, nil)
	_ = json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}

// OrderSell Place a sell order.
func OrderSell(placeRequestParams models.PlaceRequestParams) models.PlaceReturn {
	placeReturn := models.PlaceReturn{}

	mapParams := make(map[string]string)
	mapParams["volume"] = placeRequestParams.Volume
	if 0 < len(placeRequestParams.Price) {
		mapParams["price"] = placeRequestParams.Price
	}
	mapParams["pair"] = placeRequestParams.Pair

	strRequest := "/x/v1/order/sell"
	jsonPlaceReturn := utils.ApiKeyPost(mapParams, strRequest, nil)
	_ = json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}

// OrderCancelOne cancel a order by order ID. kind: 1 buy, 2 sell
func OrderCancelOne(strOrderID, kind string) models.Response {
	placeReturn := models.Response{}

	mapParams := make(map[string]string)
	mapParams["id"] = strOrderID
	mapParams["kind"] = kind

	strRequest := "/x/v1/order/cancelOne"
	jsonPlaceReturn := utils.ApiKeyPost(mapParams, strRequest, nil)
	_ = json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}

// OrderCancelByPair cancel all orders by pair.
func OrderCancelByPair(pair string) models.Response {
	placeReturn := models.Response{}

	mapParams := make(map[string]string)
	mapParams["pair"] = pair

	strRequest := "/x/v1/order/cancelByPair"
	urlValues := url.Values{}
	urlValues.Set("pair", pair)

	jsonPlaceReturn := utils.ApiKeyPost(mapParams, strRequest, urlValues)
	_ = json.Unmarshal([]byte(jsonPlaceReturn), &placeReturn)

	return placeReturn
}
