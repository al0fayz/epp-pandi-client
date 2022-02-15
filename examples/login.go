package examples

import (
	"crypto/tls"
	"epp-pandi-client/epp"
	"fmt"
	"net"
)

func Login() {
	host := HOST
	port := PORT
	url := net.JoinHostPort(host, port)
	client := &epp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	_, err := client.Connect(url)
	if err != nil {
		fmt.Println(err)
	}
	//login
	username := USERNAME
	password := PASSWORD
	login, err := client.Login(username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(login))
}
