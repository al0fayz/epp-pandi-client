package main

import (
	"crypto/tls"
	"epp-pandi-client/epp"
	"fmt"
	"net"
)

func main() {
	host := "epp-sandbox.pandi.id"
	port := "700"
	url := net.JoinHostPort(host, port)
	//use certificate

	client := &epp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	_, err := client.Connect(url)
	if err != nil {
		fmt.Println(err)
	}
	username := "H5364127"
	password := "cisauk123"
	login, err := client.Login(username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(login))
}
