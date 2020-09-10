package rgnets

import (
	"crypto/tls"
	"encoding/json"
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

// DhcpLease ...
type DhcpLease struct {
	ID       int64  `json:"id"`
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	MacAddr  string `json:"mac"`
}

// DhcpLeaseOptions ...
type DhcpLeaseOptions struct {
	Hostname string
	IP       string
}

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
