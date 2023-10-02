package cnn

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MauCastillo/alana/shared/cnn/models"
	"github.com/MauCastillo/alana/shared/request"
)

const (
	cnnAPI        = "https://production.dataviz.cnn.io/index/fearandgreed/graphdata/%s"
	timeFormatter = "2006-01-02"
)

var (
	lastTime        = ""
	dateLastRefresh = time.Now().UTC()
)

type FearAndGreedCNN struct {
	apiResponse *models.APIResponse
}

func NewFearAndGreedCNN() (*FearAndGreedCNN, error) {
	cnnResponse, err := requestCNNAPI()
	if err != nil {
		return nil, err
	}

	return &FearAndGreedCNN{apiResponse: cnnResponse}, nil
}

func (f *FearAndGreedCNN) Refresh() error {
	dateNow := time.Now().UTC()
	diff := dateNow.Sub(dateLastRefresh)

	oneHour := time.Second
	if diff < oneHour {
		return nil
	}

	cnnResponse, err := requestCNNAPI()
	if err != nil {
		return err
	}

	f.apiResponse = cnnResponse

	return nil
}

func requestCNNAPI() (*models.APIResponse, error) {

	var response models.APIResponse

	body, err := Refresh()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	dateLastRefresh = time.Now().UTC()

	return &response, nil
}

func Refresh() ([]byte, error) {
	client, err := request.NewHTTPClient()

	if err != nil {
		return []byte{}, err
	}

	header := []request.Header{{Key: "User-Agent", Value: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0"},
		{Key: "Accept", Value: "*/*"},
		{Key: "Accept-Language", Value: "en-US,en;q=0.5"},
		{Key: "Sec-Fetch-Dest", Value: "empty"},
		{Key: "Sec-Fetch-Mode", Value: "cors"},
		{Key: "Sec-Fetch-Site", Value: "cross-site"},
	}

	now := time.Now()

	lastTime = now.Format(timeFormatter)

	apiURL := fmt.Sprintf(cnnAPI, lastTime)
	body, err := client.GetwithHeaders(apiURL, header)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}

func (f *FearAndGreedCNN) Get() *models.APIResponse {
	now := time.Now()

	if lastTime != now.Format(timeFormatter) {
		cnnResponse, err := requestCNNAPI()
		if err != nil {
			return f.apiResponse
		}

		f.apiResponse = cnnResponse
	}

	return f.apiResponse
}
