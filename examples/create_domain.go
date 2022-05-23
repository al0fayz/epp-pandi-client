package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/epp"
	"github.com/al0fayz/epp-pandi-client/frames"
)

func CreateDomain() {
	domainFrame := frames.DomainCreateType{}
	domainFrame.SetDomain("bejo1.my.id")
	period := frames.Period{
		Value: 1,   //int
		Unit:  "y", //y (year) or m (month)
	}
	domainFrame.SetPeriod(period)
	//add contact
	domainFrame.SetRegistrant("hello123")
	contactAdmin := frames.Contact{
		Name: "hello123", //contactId
		Type: "admin",    //admin,tech,billing
	}
	domainFrame.AddContact(contactAdmin)
	contactTech := frames.Contact{
		Name: "hello123", //contactId
		Type: "tech",     //admin,tech,billing
	}
	domainFrame.AddContact(contactTech)
	contactBilling := frames.Contact{
		Name: "hello123", //contactId
		Type: "billing",  //admin,tech,billing
	}
	domainFrame.AddContact(contactBilling)

	//optional
	domainFrame.SetAuthInfo("pandi123") //this optional
	domainFrame.AddAddr("sn1.bejak.id")
	domainFrame.AddAddr("ns2.bejak.id")
	//add ns with ip address
	hostObj := frames.HostAttribute{
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
	}
	domainFrame.AddHostObject(hostObj)
	//print xml command create domain
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
	resCreate, err := client.CreateDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response create
	//print xml
	fmt.Println(string(resCreate))

	//handle response
	response := frames.Response{
		ResultData: &frames.DomainCreateDataType{},
	}
	if err := xml.Unmarshal(resCreate, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code is", response.Result[0].Code)
	fmt.Println("Message ", response.Result[0].Message)
	CreateData := response.ResultData.(*frames.DomainCreateDataType).CreateData
	fmt.Println("Domain ", CreateData.Name)
	fmt.Println("Create date ", CreateData.CreateDate)
	fmt.Println("Expire date ", CreateData.ExpireDate)

	//close connection
	client.Conn.Close()
}
