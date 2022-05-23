package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"

	"github.com/al0fayz/epp-pandi-client/epp"
)

func TransferContact() {
	contactFrame := frames.ContactTransferType{
		Transfer: frames.ContactTransfer{
			Name: "hello123s", //contact id
			AuthInfo: frames.AuthInfo{
				Password: "pandi123",
			},
		},
	}
	//print xml command create domain
	encoded, err := epp.Encode(contactFrame, epp.ClientXMLAttributes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encoded))

	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//update contact
	resUpdate, err := client.TransferContact(&contactFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response update
	//print xml
	fmt.Println(string(resUpdate))

	response := frames.Response{
		ResultData: &frames.ContactTransferDataType{},
	}

	if err := xml.Unmarshal(resUpdate, &response); err != nil {
		fmt.Println(err)
	}

	//print
	fmt.Println("code is", response.Result[0].Code)
	fmt.Println("Message ", response.Result[0].Message)

	transferData := response.ResultData.(*frames.ContactTransferDataType).TransferData
	fmt.Println(transferData.Name)
	fmt.Println(transferData.TransferStatus)
	fmt.Println(transferData.RequestingDate)
	fmt.Println(transferData.RequestingID)

	//close connection
	client.Conn.Close()
}
