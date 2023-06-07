package models

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

var (
	ErrURLEmpty    = errors.New("url empty")
	ErrParsingBody = errors.New("parsing body")
)

func NewHTTPClient() (*HTTPClient, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	return &HTTPClient{client: &http.Client{Transport: tr}}, nil
}

func (h *HTTPClient) Get(url string) ([]byte, error) {
	resp, err := h.client.Get(url)
	if err != nil {
		return nil, ErrURLEmpty
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrParsingBody
	}

	return body, nil
}
