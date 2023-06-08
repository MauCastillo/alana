package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"strings"

	"github.com/alana/binance-api/coins"
	"github.com/alana/binance-api/intervals"
	"github.com/alana/shared/env"
	"github.com/alana/shared/request"
)

const (

	// UIKlinePath template pathKLine
	UIKlinePath = "%s/uiKlines?symbol=%s&interval=%s&limit=%d"

	maxLimit = 1000
	minLimit = 0

	openTimeField                 = 0
	openPriceField                = 1
	highPriceField                = 2
	lowPriceField                 = 3
	closePriceField               = 4
	volumeField                   = 5
	closeTimeField                = 6
	quoteAssetVolumeField         = 7
	numberTradesField             = 8
	takerBuyBaseAssetVolumeField  = 9
	takerBuyQuoteAssetVolumeField = 10
)

var (
	// APIURL Base api binace
	APIURL = env.GetString("API_BINANCE_URL", "https://api.binance.com/api/v3")
	// ErrLimitGreater limit greater than 1000
	ErrLimitGreater = errors.New("limit greater than 1000")
	// ErrLimitLess limit less than 1
	ErrLimitLess = errors.New("limit lest than 1")
)

type UIKline struct {
	URL     string
	BodyRaw [][]interface{}
	Body    []Kline
}

type Kline struct {
	OpenTime                 int64  `json:"open_time"`
	OpenPrice                string `json:"open_price"`
	HighPrice                string `json:"high_price"`
	LowPrice                 string `json:"low_price"`
	ClosePrice               string `json:"close_price"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"close_time"`
	QuoteAssetVolume         string `json:"quoteAsset_volume"`
	NumberTrades             int    `json:"number_trades"`
	TakerBuyBaseAssetVolume  string `json:"taker_buy_base_asset_volume"`
	TakerBuyQuoteAssetVolume string `json:"taker_buy_quote_asset_volume"`
}

func NewUIKline(symbol *coins.Coin, internal *interval.Interval, limit int64) (*UIKline, error) {

	if limit > maxLimit {
		return nil, ErrLimitGreater
	}

	if limit < minLimit {
		return nil, ErrLimitLess
	}

	requestURL := fmt.Sprintf(UIKlinePath, APIURL, symbol.Value, internal.Value, limit)

	client, err := request.NewHTTPClient()
	if err != nil {
		return nil, err
	}

	body, err := client.Get(requestURL)
	if err != nil {
		return nil, err
	}

	request, err := convertToArray(body)
	if err != nil {
		return nil, err
	}

	return &UIKline{
		URL:     requestURL,
		BodyRaw: request}, nil
}

func convertToArray(body []byte) ([][]interface{}, error) {
	var arr [][]interface{}
	err := json.Unmarshal([]byte(body), &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}

func main() {
	objectKline, err := NewUIKline(coins.EthBusd, interval.ThirtyMinutes, 2)
	if err != nil{
		fmt.Print(err.Error())
	}

	for index, klineRaw := range objectKline.BodyRaw{
		fmt.Println(index)
		for _, value := range klineRaw{
			fmt.Println(value)
		}
		
	}
		

	/*klines := []Kline{}

	 for index, klineRaw := range uiKline.BodyRaw{
			fmt.Print(index)
			kline := strings.Split(klineRaw, " ")
			line := Kline{
				//OpenTime:                 kline[openTimeField],
				OpenPrice:                kline[openPriceField],
				HighPrice:                kline[highPriceField],
				LowPrice:                 kline[lowPriceField],
				ClosePrice:               kline[closePriceField],
				Volume:                   kline[volumeField],
				//CloseTime:                kline[closeTimeField],
				QuoteAssetVolume:         kline[quoteAssetVolumeField],
				//NumberTrades:             kline[numberTradesField],
				TakerBuyBaseAssetVolume:  kline[takerBuyBaseAssetVolumeField],
				TakerBuyQuoteAssetVolume: kline[takerBuyQuoteAssetVolumeField]}

				klines = append(klines, line)
		} */

}
