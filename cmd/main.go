package main

import (
	"fmt"
	"os"

	"github.com/drkchiloll/rgnets"
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
	doptions := rgnets.DhcpLeaseOptions{Hostname: "ap01.rXgCluster"}
	dlease, _ := rxg.GetDhcpLease(doptions)
	fmt.Println(dlease)
}
