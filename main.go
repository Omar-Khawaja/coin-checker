package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/hashicorp/hcl"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	AccountSid string `hcl:"accountSid"`
	AuthToken  string `hcl:"authToken"`
	FromNumber string `hcl:"fromNumber"`
	ToNumber   string `hcl:"toNumber"`
}

type Crypto []struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

var msg string
var config Config

func main() {

	coin := flag.String("coin", "bitcoin", "unit of cryptocurency")
	info := flag.String("info", "price", "info level")
	flag.Parse()

	if err := getCryptoInfo(*coin, *info); err != nil {
		log.Fatal(err)
	}
}

func LoadConfiguration(file string) {
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = hcl.Decode(&config, string(configFile))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getCryptoInfo(coin, info string) error {
	LoadConfiguration("credentials.conf")

	resp, err := http.Get(fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s/", coin))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var crypto Crypto
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&crypto); err != nil {
		fmt.Println("Could not retrieve info on coin")
		return err
	}

	FieldMap := map[string]string{
		"ID":               crypto[0].ID,
		"Name":             crypto[0].Name,
		"Symbol":           crypto[0].Symbol,
		"Rank":             crypto[0].Rank,
		"PriceUsd":         crypto[0].PriceUsd,
		"PriceBtc":         crypto[0].PriceBtc,
		"Two4HVolumeUsd":   crypto[0].Two4HVolumeUsd,
		"MarketCapUsd":     crypto[0].MarketCapUsd,
		"AvailableSupply":  crypto[0].AvailableSupply,
		"TotalSupply":      crypto[0].TotalSupply,
		"MaxSupply":        crypto[0].MaxSupply,
		"PercentChange1H":  crypto[0].PercentChange1H,
		"PercentChange24H": crypto[0].PercentChange24H,
		"PercentChange7D":  crypto[0].PercentChange7D,
		"LastUpdated":      crypto[0].LastUpdated,
	}

	info_level := info
	switch info_level {
	case "price":
		msg = fmt.Sprintf("The price of %v today is $%v\n", crypto[0].Name, crypto[0].PriceUsd)
	case "all":
		var str string
		for k, v := range FieldMap {
			str += fmt.Sprintln(k, v)
		}
		msg = str
	}
	SendSMS(msg, config.AccountSid, config.AuthToken, config.ToNumber, config.FromNumber)
	return err
}
