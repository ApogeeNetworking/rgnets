package rgnets

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// Client struct used for the Connection to an RxG RESTful API
type Client struct {
	BaseURL string
	http    *http.Client
	apiKey  string
}

// New creates a reference to the Client struct
func New(host, apiKey string, ignoreSSL bool) *Client {
	return &Client{
		BaseURL: fmt.Sprintf("https://%s/admin/scaffolds", host),
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: ignoreSSL,
				},
			},
			Timeout: 8 * time.Second,
		},
		apiKey: apiKey,
	}
}
