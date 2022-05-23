package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/epp"
	"github.com/al0fayz/epp-pandi-client/frames"
)

func UpdateDomain() {
	domainFrame := frames.DomainUpdateType{}
	domainFrame.SetDomain("bejak.id")
	//change data
	domainChange := frames.DomainChange{
		Registrant: "hello123s", //if you want change registrant contact
		AuthInfo: &frames.AuthInfo{
			Password: "Helloworld",
		},
	}
	domainFrame.ChangeData(domainChange)
	//remove data
	removeData := frames.DomainAddRemove{
		//name server
		NameServer: frames.NameServer{
			HostObject: []string{
				"sn1.bejak.id",
				"sn2.bejak.id",
			},
			HostAttribute: []frames.HostAttribute{
				{
					HostName: "ns1.bejak.id",
					HostAddress: []frames.HostAddress{
						{
							Address: "127.0.0.1",
							IP:      frames.HostIPv4,
						},
						{
							Address: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
							IP:      frames.HostIPv6,
						},
					},
				},
			},
		},

		//contact remove
		Contact: []frames.Contact{
			{
				Name: "hello123",
				Type: "admin",
			},
			{
				Name: "hello123",
				Type: "tech",
			},
			{
				Name: "hello123",
				Type: "billing",
			},
		},

		//remove status
		Status: []frames.DomainStatus{
			{
				Status:           "clientDeleteProhibited",
				DomainStatusType: frames.DomainStatusClientDeleteProhibited,
			},
		},
	}
	domainFrame.RemoveData(removeData)

	//add data
	addData := frames.DomainAddRemove{
		//name server
		NameServer: frames.NameServer{
			HostObject: []string{
				"ns4.bejak.id",
				"ns3.bejak.id",
			},
			HostAttribute: []frames.HostAttribute{
				{
					HostName: "ns5.bejak.id",
					HostAddress: []frames.HostAddress{
						{
							Address: "127.0.0.1",
							IP:      frames.HostIPv4,
						},
						{
							Address: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
							IP:      frames.HostIPv6,
						},
					},
				},
			},
		},

		//contact remove
		Contact: []frames.Contact{
			{
				Name: "hello123s",
				Type: "admin",
			},
			{
				Name: "hello123s",
				Type: "tech",
			},
			{
				Name: "hello123s",
				Type: "billing",
			},
		},

		//remove status
		Status: []frames.DomainStatus{
			{
				Status:           "clientTransferProhibited",
				DomainStatusType: frames.DomainStatusClientTransferProhibited,
			},
		},
	}
	domainFrame.AddData(addData)
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
