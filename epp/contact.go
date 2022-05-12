package epp

import (
	"epp-pandi-client/frames"
)

//check
func (c *Client) CheckContact(contact *frames.ContactCheckType) ([]byte, error) {
	encoded, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}

//info
func (c *Client) InfoContact(contact *frames.ContactInfoType) ([]byte, error) {
	encoded, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}

//create
func (c *Client) CreateContact(contact *frames.ContactCreateType) ([]byte, error) {
	encode, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

//update
func (c *Client) UpdateContact(contact *frames.ContactUpdateType) ([]byte, error) {
	encode, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

///delete
func (c *Client) DeleteContact(contact *frames.ContactDeleteType) ([]byte, error) {
	encode, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}

//transfer
func (c *Client) TransferContact(contact *frames.ContactTransferType) ([]byte, error) {
	encode, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
