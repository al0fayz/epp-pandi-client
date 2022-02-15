package epp

import (
	"epp-pandi-client/frames"
	"fmt"
)

//check
func (c *Client) CheckHost(host *frames.HostCheckType) ([]byte, error) {
	encoded, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}

//info
func (c *Client) InfoHost(host *frames.HostInfoType) ([]byte, error) {
	encode, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

//create
func (c *Client) CreateHost(host *frames.HostCreateType) ([]byte, error) {
	encode, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

//update
func UpdateHost(host *frames.HostUpdateType) {
	encode, err := Encode(host, ClientXMLAttributes())
	fmt.Println(string(encode), err)
}

//delete
func (c *Client) DeleteHost(host *frames.HostDeleteType) ([]byte, error) {
	encode, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
