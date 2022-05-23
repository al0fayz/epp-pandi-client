package examples

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"github.com/al0fayz/epp-pandi-client/epp"
)

func ExampleLogin() {
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

	//logout
	logout, err := client.Logout()
	if err != nil {
		fmt.Println(err)
	}
	//response logout
	fmt.Println(string(logout))
}
func Login() (*epp.Client, error) {
	host := HOST
	port := PORT
	url := net.JoinHostPort(host, port)
	//use certificate
	certPath := CERTPATH
	keyPath := KEYPATH
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	username := USERNAME
	password := PASSWORD
	_, err = client.Login(username, password)
	if err != nil {
		return nil, err
	}
	return client, nil
}
