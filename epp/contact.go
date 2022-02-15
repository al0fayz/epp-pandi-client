package epp

import (
	"epp-pandi-client/frames"
)

func (c *Client) CheckContact(contact *frames.ContactCheckType) ([]byte, error) {
	encoded, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
func (c *Client) InfoContact(contact *frames.ContactInfoType) ([]byte, error) {
	encoded, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
func CreateContact() ([]byte, error) {
	return nil, nil
}
