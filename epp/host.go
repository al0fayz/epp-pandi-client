package epp

import (
	"epp-pandi-client/frames"
)

func (c *Client) CheckHost(host *frames.HostCheckType) ([]byte, error) {
	encoded, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
func (c *Client) InfoHost(host *frames.HostInfoType) ([]byte, error) {
	encode, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
