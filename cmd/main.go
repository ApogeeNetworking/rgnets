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
	// leaseID := "1048577"
	// rxg.DhcpRelease(leaseID)
	// doptions := rgnets.DhcpLeaseOptions{IP: "ipAddr"}
	// doptions := rgnets.DhcpLeaseOptions{Hostname: "ap01.rXgCluster"}
	// dlease, _ := rxg.GetDhcpLease(doptions)
	// fmt.Println(dlease)
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
