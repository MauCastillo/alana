package request

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/alana/shared/env"
)

type HTTPClient struct {
	client *http.Client
}

var (
	ErrURLEmpty    = errors.New("url empty")
	ErrParsingBody = errors.New("parsing body")

	timeoutConnections = env.GetInt64("TIMEOUT", 10)
	maxLimitConnections = env.GetInt64("MAX_LIMIT_CONNECTIONS", 10)
	DisableCompression = env.GetBool("DISABLE_COMPRESSION", true)
)

func NewHTTPClient() (*HTTPClient, error) {
	tr := &http.Transport{
		MaxIdleConns:       int(maxLimitConnections),
		IdleConnTimeout:    time.Duration(timeoutConnections) * time.Second,
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
