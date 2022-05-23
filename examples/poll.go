package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"

	"github.com/al0fayz/epp-pandi-client/epp"
)

func PollRequest() {
	frame := frames.Poll{
		Poll: frames.PollCommand{
			Operation: frames.PollOperationRequest, //poll request
		},
	}
	//print xml command
	encoded, err := epp.Encode(frame, epp.ClientXMLAttributes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encoded))

	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//poll
	poll, err := client.Poll(&frame)
	if err != nil {
		fmt.Println(err)
	}
	//response error
	//print xml
	fmt.Println(string(poll))
	response := frames.Response{
		ResultData: &frames.DomainTransferDataType{},
	}
	if err := xml.Unmarshal(poll, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Code", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)

	//trnData
	trnData := response.ResultData.(*frames.DomainTransferDataType).TransferData
	fmt.Println("Domain ", trnData.Name)
	fmt.Println("Transfer Status ", trnData.TransferStatus)
	fmt.Println("Request Id ", trnData.RequestingID)
	fmt.Println("Request Date", trnData.RequestingDate)
	fmt.Println("Acting Id", trnData.ActingID)

	//messageQ
	fmt.Println("Message Q", response.MessageQ.Message)
	fmt.Println("Count ", response.MessageQ.Count)
	fmt.Println("Id", response.MessageQ.ID)

	//close connection
	client.Conn.Close()
}
