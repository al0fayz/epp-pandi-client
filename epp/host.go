package epp

import (
	"epp-pandi-client/types"
)

func (c *Client) CheckHost(hosts []string) ([]byte, error) {
	host := types.HostCheckType{
		Check: types.HostCheck{
			Names: hosts,
		},
	}
	encoded, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
func (c *Client) InfoHost(name string) ([]byte, error) {
	host := types.HostInfoType{
		Info: types.HostInfo{
			Name: name,
		},
	}
	encode, err := Encode(host, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
