package epp

import "epp-pandi-client/frames"

func (c *Client) Poll(data *frames.Poll) ([]byte, error) {
	encode, err := Encode(data, ClientXMLAttributes())
	if err != nil {
		return nil, err
	}
	return c.Send(encode)
}
