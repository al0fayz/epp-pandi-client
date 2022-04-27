package examples

import (
	"encoding/xml"
	"epp-pandi-client/frames"
	"fmt"
)

func DeleteHost() {
	hostFrame := frames.HostDeleteType{}
	hostFrame.SetHost("ns3.bejak.id")
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	resDelete, err := client.DeleteHost(&hostFrame)
	if err != nil {
		fmt.Println(err)
	}
	//print xml
	fmt.Println(string(resDelete))
	response := frames.Response{}
	if err := xml.Unmarshal(resDelete, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Code is ", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)

	client.Conn.Close() //close connection
}
