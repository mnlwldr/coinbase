package coinbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	COINBASE_API_URL = "https://api.coinbase.com/v2/prices"
	TYPE             = "spot"
)

type Response struct {
	Amount   float64
	Base     string
	Currency string
}

func buildUrl(coin string) string {
	return fmt.Sprintf("%s/%s/%s", COINBASE_API_URL, coin, TYPE)
}

func buildUrlWithDate(coin, date string) string {
	return fmt.Sprintf("%s/%s/%s?date=%s", COINBASE_API_URL, coin, TYPE, date)
}

func get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

// The coin abbreviation is passed as parameter "coin".
// The coin abbreviation is a currency pair. For example: "SHIB-EUR" or "BTC-USD"
// The date must be in the format YYYY-MM-DD
func GetWithDate(coin, date string) (Response, error) {
	url := buildUrlWithDate(coin, date)
	return buildResponse(url)
}

// The coin abbreviation is passed as parameter "coin".
// The coin abbreviation is a currency pair. For example: "SHIB-EUR" or "BTC-USD"
func Get(coin string) (Response, error) {
	url := buildUrl(coin)
	return buildResponse(url)
}

func buildResponse(url string) (Response, error) {
	response, err := get(url)
	if err != nil {
		return Response{}, err
	}

	var r interface{}
	if err := json.Unmarshal(response, &r); err != nil {
		return Response{}, err
	}

	f := r.(map[string]interface{})["data"].(map[string]interface{})["amount"].(string)
	amount, err := strconv.ParseFloat(f, 64)

	if err != nil {
		return Response{}, err
	}
	base := r.(map[string]interface{})["data"].(map[string]interface{})["base"].(string)
	currency := r.(map[string]interface{})["data"].(map[string]interface{})["currency"].(string)

	return Response{
		Amount:   amount,
		Base:     base,
		Currency: currency,
	}, nil
}
