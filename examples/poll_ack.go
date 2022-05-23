package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"

	"github.com/al0fayz/epp-pandi-client/epp"
)

func PollAck() {
	messageId := "1548"
	frame := frames.Poll{
		Poll: frames.PollCommand{
			Operation: frames.PollOperationAcknowledge, //ack
			MessageID: &messageId,
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
	response := frames.Response{}

	if err := xml.Unmarshal(poll, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Code", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)

	//close connection
	client.Conn.Close()
}
