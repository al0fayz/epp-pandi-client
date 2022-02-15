package epp

import (
	"epp-pandi-client/types"
)

func (c *Client) CheckDomain(domains []string) ([]byte, error) {
	domain := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: domains,
		},
	}
	encoded, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
func (c *Client) InfoDomain(name string) ([]byte, error) {
	domain := types.DomainInfoType{
		Info: types.DomainInfo{
			Name: types.DomainInfoName{
				Name: name,
			},
		},
	}
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
