package http

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// Get performs an HTTP GET request and returns the response
func Get(url string) (*http.Response, error) {
	return httpClient.Get(url)
}
