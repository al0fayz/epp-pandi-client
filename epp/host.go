package epp

import (
	"github.com/al0fayz/epp-pandi-client/frames"
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
func (c *Client) UpdateHost(host *frames.HostUpdateType) ([]byte, error) {
	encode, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

//delete
func (c *Client) DeleteHost(host *frames.HostDeleteType) ([]byte, error) {
	encode, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
