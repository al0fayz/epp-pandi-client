package epp

import (
	"github.com/al0fayz/epp-pandi-client/frames"
)

/**
Created by ahmad ropai <akhmadrofai@gmail.com>
*/
//check
func (c *Client) CheckDomain(domain *frames.DomainCheckType) ([]byte, error) {
	encoded, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}

// info
func (c *Client) InfoDomain(domain *frames.DomainInfoType) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

// create
func (c *Client) CreateDomain(domain *frames.DomainCreateType) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

// update
func (c *Client) UpdateDomain(domain *frames.DomainUpdateType) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

// renew
func (c *Client) RenewDomain(domain *frames.DomainRenewType) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

// delete
func (c *Client) DeleteDomain(domain *frames.DomainDeleteType) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(encode))
	return c.Send(encode)
}

// transfer
func (c *Client) TransferDomain(domain *frames.DomainTransferTypeCommand) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(encode))
	return c.Send(encode)
}

// restore
func (c *Client) RestoreDomain(domain *frames.DomainRestoreType) ([]byte, error) {
	encode, err := Encode(domain, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(encode))
	return c.Send(encode)
}
