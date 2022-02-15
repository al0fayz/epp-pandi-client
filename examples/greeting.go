package examples

import (
	"crypto/tls"
	"epp-pandi-client/epp"
	"fmt"
	"net"
)

func Greeting() {
	host := HOST
	port := PORT
	url := net.JoinHostPort(host, port)
	client := &epp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	greeting, err := client.Connect(url)
	if err != nil {
		fmt.Println(err)
	}
	//response greeting
	fmt.Println(string(greeting))
	client.Conn.Close()
}
