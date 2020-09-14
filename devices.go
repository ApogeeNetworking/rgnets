package rgnets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetDevice ...
func (c *Client) GetDevice(options GetDeviceOptions) ([]Device, error) {
	var devices []Device
	ep := fmt.Sprintf("/devices/index.json?api_key=%s", c.apiKey)

	switch {
	case options.AccountID != 0:
		ep += fmt.Sprintf("&account_id=%v", options.AccountID)
	case options.DeviceID != 0:
		ep += fmt.Sprintf("&id=%v", options.DeviceID)
	case options.MacAddr != "":
		ep += fmt.Sprintf("&mac=%s", options.MacAddr)
	}
	req, err := http.NewRequest("GET", c.BaseURL+ep, nil)
	if err != nil {
		return devices, err
	}
	res, err := c.http.Do(req)
	if err != nil {
		return devices, err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&devices)
	return devices, nil
}
