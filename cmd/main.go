package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ApogeeNetworking/rgnets"
	"github.com/subosito/gotenv"
)

var host, apiKey string

func init() {
	gotenv.Load()
	host = os.Getenv("HOST")
	apiKey = os.Getenv("API_KEY")
}

func main() {
	rxg := rgnets.New(host, apiKey, true)
	accountOptions := rgnets.GetAccountOptions{
		ID: 88,
	}
	account, _ := rxg.GetAccount(accountOptions)
	fmt.Println(account)
}

func dhcpOps(rxg *rgnets.Client) {
	dlops := rgnets.DhcpLeaseOptions{
		IP: "ipAddr",
		// Hostname: "ap01.rXgCluster",
	}
	dlease, _ := rxg.GetDhcpLease(dlops)
	fmt.Println(dlease)

	// Release the DHCP Lease
	rxg.ReleaseDHCP(dlease.ID)
}

func getDevice(rxg *rgnets.Client) {
	dops := rgnets.GetDeviceOptions{
		// MacAddr: "78:45:61:c2:93:9f",
		// DeviceID: 111,
		AccountID: 88,
	}
	devs, e := rxg.GetDevice(dops)
	if e != nil {
		log.Fatalf("%v", e)
	}
	fmt.Println(devs)
}
