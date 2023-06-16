package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseUrl = "https://graph.facebook.com"
	Version = "v17.0"
	Endpoint = "likes"
	AccessToken = "Your Access Token"
)

func main() {
	pageId := "101372079670424_118057034648183"
	url := fmt.Sprintf("%v/%v/%s/%v?since=%v&until=%v&access_token=%v", BaseUrl, Version, pageId, Endpoint, "2023-06-15", "2023-06-16", AccessToken);
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}
