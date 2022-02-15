package epp

import (
	"epp-pandi-client/types"
)

func (c *Client) CheckContact(contacts []string) ([]byte, error) {
	contact := types.ContactCheckType{
		Check: types.ContactCheck{
			Names: contacts,
		},
	}
	encoded, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encoded)
}
