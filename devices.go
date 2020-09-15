package rgnets

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GetDevice ...
func (c *Client) GetDevice(options GetDeviceOptions) ([]Device, error) {
	var devices []Device
	ep := "/devices/index.json"

	req, err := http.NewRequest("GET", c.BaseURL+ep, nil)
	if err != nil {
		return devices, err
	}
	// Create a Proper QueryString for REQ
	qs := req.URL.Query()
	qs.Add("api_key", c.apiKey)
	switch {
	case options.AccountID != 0:
		qs.Add("account_id", strconv.Itoa(options.AccountID))
	case options.DeviceID != 0:
		qs.Add("id", strconv.Itoa(options.DeviceID))
	case options.MacAddr != "":
		qs.Add("mac", options.MacAddr)
	}
	// Add QueryString Encoded properly
	req.URL.RawQuery = qs.Encode()
	res, err := c.http.Do(req)
	if err != nil {
		return devices, err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&devices)
	return devices, nil
}
