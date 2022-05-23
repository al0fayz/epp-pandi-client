package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/epp"
	"github.com/al0fayz/epp-pandi-client/frames"
)

func UpdateDomainDnssec() {
	domainFrame := frames.DomainUpdateType{}
	domainFrame.SetDomain("bejo5.my.id")

	dnsSec := frames.DNSSECExtensionUpdateType{
		Update: frames.DNSSECExtensionUpdate{
			Remove: frames.DNSSECRemove{
				All: false, //change true if you want remove all dnssec
			},
			Add: frames.DNSSECOrKeyData{
				DNSSECData: []frames.DNSSEC{
					{
						KeyTag:     45114,
						Algorithm:  13,
						DigestType: 1,
						Digest:     "4030544B4739A9A090D94FEED88FEE9157CE792A",
					},
					{
						KeyTag:     45114,
						Algorithm:  13,
						DigestType: 2,
						Digest:     "4978CA3D92CD618CC8079A6262C0F3C6C27F60F5CDF4DD366F18ABE1D10C010C",
					},
				},
			},
		},
	}
	domainFrame.ChangeDnssec(dnsSec)
	//print xml command update contact
	encoded, err := epp.Encode(domainFrame, epp.ClientXMLAttributes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encoded))

	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//create domain
	resCreate, err := client.UpdateDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response create
	//print xml
	fmt.Println(string(resCreate))

	//handle response
	response := frames.Response{}

	if err := xml.Unmarshal(resCreate, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code is", response.Result[0].Code)
	fmt.Println("Message ", response.Result[0].Message)

	//close connection
	client.Conn.Close()
}
