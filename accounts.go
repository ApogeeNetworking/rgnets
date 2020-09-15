package rgnets

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetAccount ...
func (c *Client) GetAccount(options GetAccountOptions) ([]Account, error) {
	var accounts []Account
	ep := "/accounts/index.json"

	req, err := http.NewRequest("GET", c.BaseURL+ep, nil)
	if err != nil {
		return accounts, err
	}
	// Create a Proper QueryString for the Request
	qs := req.URL.Query()
	qs.Add("api_key", c.apiKey)
	switch {
	case options.ID != 0:
		qs.Add("id", strconv.Itoa(options.ID))
	case options.Username != "":
		qs.Add("login", options.Username)
	}
	req.URL.RawQuery = qs.Encode()
	res, err := c.http.Do(req)
	if err != nil {
		return accounts, err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&accounts)
	return accounts, nil
}
