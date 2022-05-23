package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"

	"github.com/al0fayz/epp-pandi-client/epp"
)

func TransferDomainQuery() {
	domainFrame := frames.DomainTransferTypeCommand{}
	domainFrame.SetDomain("createdomaintest.id")
	domainFrame.SetAuthInfo("aVfJjXcx")
	domainFrame.SetOpertaion("query") //query, request
	period := frames.Period{
		Value: 1,
		Unit:  "y",
	}
	domainFrame.SetPeriod(period)

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
	resCreate, err := client.TransferDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response create
	//print xml
	fmt.Println(string(resCreate))

	//handle response
	response := frames.Response{
		ResultData: &frames.DomainTransferDataType{},
	}

	if err := xml.Unmarshal(resCreate, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code is", response.Result[0].Code)
	fmt.Println("Message ", response.Result[0].Message)

	//trnData
	trnData := response.ResultData.(*frames.DomainTransferDataType).TransferData
	fmt.Println("Domain ", trnData.Name)
	fmt.Println("Transfer Status ", trnData.TransferStatus)
	fmt.Println("Request Id ", trnData.RequestingID)
	fmt.Println("Request Date", trnData.RequestingDate)
	fmt.Println("Acting Id", trnData.ActingID)

	//close connection
	client.Conn.Close()
}
