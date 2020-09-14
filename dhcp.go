package rgnets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetDhcpLease ...
func (c *Client) GetDhcpLease(options DhcpLeaseOptions) (DhcpLease, error) {
	endpoint := "/dhcp_leases/index.json"
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	q := req.URL.Query()
	q.Add("api_key", c.apiKey)
	// Add Query String for Weird Hostnames to be Encoded Properly
	q.Add("hostname", options.Hostname)
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return DhcpLease{}, err
	}
	res, err := c.http.Do(req)
	if err != nil {
		return DhcpLease{}, err
	}
	defer res.Body.Close()
	var leases []DhcpLease
	json.NewDecoder(res.Body).Decode(&leases)
	if len(leases) == 0 {
		return DhcpLease{}, nil
	}
	return leases[0], nil
}

// ReleaseDHCP ...
func (c *Client) ReleaseDHCP(leaseID int64) error {
	endpoint := fmt.Sprintf("/dhcp_leases/destroy/%v.xml?api_key=%s", leaseID, c.apiKey)
	req, err := http.NewRequest("POST", c.BaseURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	_, err = c.http.Do(req)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
