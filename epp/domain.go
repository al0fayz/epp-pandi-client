package epp

import (
	"epp-pandi-client/frames"
)

func (c *Client) CheckDomain(domain *frames.DomainCheckType) ([]byte, error) {
	encoded, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
func (c *Client) InfoDomain(domain *frames.DomainInfoType) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
