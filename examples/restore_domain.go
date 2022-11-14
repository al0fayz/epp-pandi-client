package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/epp"
	"github.com/al0fayz/epp-pandi-client/frames"
)

func RestoreDomain() {
	domainFrame := frames.DomainRestoreType{}
	domainFrame.SetDomain("bejo1.my.id")
	domainFrame.Extension = &frames.RgpUpdateType{
		Rgp: frames.RgpType{
			Restore: frames.Restore{
				Op: "request",
			},
		},
	}

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
	//restore domain
	resTransfer, err := client.RestoreDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response error
	//print xml
	fmt.Println(string(resTransfer))
	response := frames.Response{
		ResultData: &frames.DomainRenewDataType{},
	}
	if err := xml.Unmarshal(resTransfer, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Code", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)

	//close connection
	client.Conn.Close()
}
