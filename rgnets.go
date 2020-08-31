package rgnets

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
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
	endpoint := fmt.Sprintf("/dhcp_leases/index.json?api_key=%s", c.apiKey)
	switch {
	case options.Hostname != "":
		endpoint += "&hostname=" + options.Hostname
	case options.IP != "":
		endpoint += "&ip=" + options.IP
	}
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
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

// DhcpRelease ...
func (c *Client) DhcpRelease(leaseID string) {
	endpoint := fmt.Sprintf("/dhcp_leases/destroy/%s.xml?api_key=%s", leaseID, c.apiKey)
	req, err := http.NewRequest("POST", c.BaseURL+endpoint, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	res, err := c.http.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(res)
}
