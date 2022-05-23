package epp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/al0fayz/epp-pandi-client/frames"
)

type Client struct {
	//this is connection from epp
	Conn      net.Conn
	TLSConfig *tls.Config
}

func (client *Client) Connect(url string) ([]byte, error) {
	//10 second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure cancel is always called
	d := tls.Dialer{
		Config: client.TLSConfig,
	}
	//connect to server
	conn, err := d.DialContext(ctx, "tcp", url)
	if err != nil {
		return nil, err
	}
	greeting, err := ReadMessage(conn)
	if err != nil {
		return nil, err
	}
	client.Conn = conn
	return greeting, err
}

// Send will send data to the server.
func (c *Client) Send(data []byte) ([]byte, error) {
	err := WriteMessage(c.Conn, data)
	if err != nil {
		_ = c.Conn.Close()

		return nil, err
	}

	_ = c.Conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	msg, err := ReadMessage(c.Conn)
	if err != nil {
		_ = c.Conn.Close()

		return nil, err
	}

	return msg, nil
}

// Login will perform a login to an EPP server.
func (c *Client) Login(username, password string) ([]byte, error) {
	login := frames.Login{
		ClientID: username,
		Password: password,
		Options: frames.LoginOptions{
			Version:  "1.0",
			Language: "en",
		},
		Services: frames.LoginServices{
			ObjectURI: []string{
				"urn:ietf:params:xml:ns:domain-1.0",
				"urn:ietf:params:xml:ns:contact-1.0",
				"urn:ietf:params:xml:ns:host-1.0",
			},
			ServiceExtension: frames.LoginServiceExtension{
				ExtensionURI: []string{
					"urn:ietf:params:xml:ns:secDNS-1.0",
					"urn:ietf:params:xml:ns:secDNS-1.1",
				},
			},
		},
	}

	encoded, err := Encode(login, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
func (c *Client) Logout() ([]byte, error) {
	logout := frames.Logout{}
	encoded, err := Encode(logout, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	fmt.Println(string(encoded))
	return c.Send(encoded)
}
