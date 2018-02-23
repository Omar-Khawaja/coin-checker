package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SendSMS(stats, accountSid, authToken, toNumber string, fromNumber string) {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", toNumber)
	msgData.Set("From", fromNumber)
	msgData.Set("Body", stats)
	msgDataReader := strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlStr, msgDataReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&data); err == nil {
			fmt.Println(data["sid"])
		} 
	} else {
		fmt.Println(resp.Status)
	}
}
