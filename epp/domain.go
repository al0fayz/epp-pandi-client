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
