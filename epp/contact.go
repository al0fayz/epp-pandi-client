package epp

import (
	"epp-pandi-client/types"
	"fmt"
)

func (c *Client) CheckContact() ([]byte, error) {
	contact := types.ContactCheckType{
		types.ContactCheck{
			Names: []string{
				"hello",
				"test",
			},
		},
	}
	encoded, err := Encode(contact, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	fmt.Println(string(encoded))
	return nil, nil
}
