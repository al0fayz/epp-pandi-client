package examples

import (
	"crypto/tls"
	"epp-pandi-client/epp"
	"fmt"
	"log"
	"net"
)

func Login() {
	host := HOST
	port := PORT
	url := net.JoinHostPort(host, port)
	//use certificate
	certPath := CERTPATH
	keyPath := KEYPATH
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal(err)
	}
	client := &epp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{cert},
		},
	}
	//connect to epp server
	_, err = client.Connect(url)
	if err != nil {
		fmt.Println(err)
	}
	username := USERNAME
	password := PASSWORD
	login, err := client.Login(username, password)
	if err != nil {
		fmt.Println(err)
	}
	//response login
	fmt.Println(string(login))
	defer client.Conn.Close()
}
