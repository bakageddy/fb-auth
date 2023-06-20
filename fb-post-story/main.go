// Shoot the killi
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	// AppID for fb
	AppID = "284950570575577"
	// AccessToken for fb
	AccessToken = "EAAEDKTmJptkBAMUAe4wwS0m1MXMTZCZB7NXlSSIK2pGhOsemZCKTyrj2vuTGqvQABubd09WrqvjmQ2GZB3jjmMh4oZC66TWlktWAnZCAeWafHlWgrAJefMtR7gvBlymPoGGWpGBCVTV2PoISezgzC9JA4idqKHeb6i6oO9ucsTf74YoSpSiMZCDf9GDCy8C3PqMB2JZCrp8wAmrqByguaslB" 
	// BaseURL for api
	BaseURL = "https://graph.facebook.com/v17.0"
	// UploadURL for uploading urls
	UploadURL = "https://rupload.facebook.com/video-upload/v17.0"
	// EndPoint for api
	EndPoint = "videos"
	// VideoURL rickroll video
	VideoURL = "https://youtu.be/xvFZjo5PgG0"
	// PageID to post
	PageID = 101372079670424
)

// FacebookVideosResponse i dont wanna explain
type FacebookVideosResponse struct {
	ID string `json:"id"`
}

func main() {
	var result FacebookVideosResponse
	auth := "OAuth " + AccessToken
	apiURL := fmt.Sprintf("%v/%v/%v", BaseURL, PageID, EndPoint)

	params := url.Values{}
	params.Add("access_token", AccessToken)
	params.Add("file_url", VideoURL)
	params.Add("description", "test")

	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(params.Encode()))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer res.Body.Close()

	_ = json.NewDecoder(res.Body).Decode(&result)
	uploadURL := fmt.Sprintf("%v/%v", UploadURL, result.ID)

	uploadReq, err := http.NewRequest(http.MethodPost, uploadURL ,nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	uploadReq.Header.Add("Authorization", auth)
	uploadReq.Header.Add("Content-Type", "application/json")
	uploadReq.Header.Add("video_url", VideoURL)

	uploadRes, err := http.DefaultClient.Do(uploadReq)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer uploadRes.Body.Close()
	body, _ := ioutil.ReadAll(uploadRes.Body)
	fmt.Println(string(body))
}
