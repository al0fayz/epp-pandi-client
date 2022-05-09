package examples

import (
	"encoding/xml"
	"epp-pandi-client/frames"
	"fmt"
)

func DeleteDomain() {
	domainFrame := frames.DomainDeleteType{}
	domainFrame.SetDomain("bejo.my.id")
	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//delete domain
	resDelete, err := client.DeleteDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response delete
	//print xml
	fmt.Println(string(resDelete))
	response := frames.Response{}
	if err := xml.Unmarshal(resDelete, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Code", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)

	//close connection
	client.Conn.Close()
}
