package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type VonageResponse struct {
	Messages []struct {
		Status    string `json:"status"`
		ErrorText string `json:"error-text"`
	} `json:"messages"`
}

func main() {
	apiKey := "ec047236"
	apiSecret := "CTFPhvpIR65EQQ31"
	apiPath := "https://rest.nexmo.com/sms/json"

	from := "Vonage APIs"
	to := "919496360305"

	message := "Hello Vonage SMS API"

	// Request body
	body := url.Values{
		"from":			{from},
		"to":         	{to},
		"text":       	{message},
		"api_key":    	{apiKey},
		"api_secret": 	{apiSecret},
	}

	// Create a HTTP post request
	client := &http.Client{}
	r, err := http.NewRequest("POST", apiPath, strings.NewReader(body.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	resp := &VonageResponse{}
	derr := json.NewDecoder(res.Body).Decode(resp)
	if derr != nil {
		fmt.Println(err)
		return
	}

	if len(resp.Messages) <= 0 {
		fmt.Println("Vonage error: Internal Error")
		return
	}

	// A status of zero indicates success; a non-zero value means something went wrong.
	if resp.Messages[0].Status != "0" {
		fmt.Errorf("Vonage error: %v (status: %v)", resp.Messages[0].ErrorText, resp.Messages[0].Status)
		return
	}

	fmt.Println("SMS sent successfully.")
}