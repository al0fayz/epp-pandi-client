package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"
)

func DeleteContact() {
	contactFrame := frames.ContactDeleteType{}
	contactFrame.SetContactId("hello123s")
	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//delete
	resDelete, err := client.DeleteContact(&contactFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response delete
	//print xml
	// fmt.Println(string(resDelete))

	response := frames.Response{}
	if err := xml.Unmarshal(resDelete, &response); err != nil {
		fmt.Println(err)
	}
	//print
	fmt.Println("code is", response.Result[0].Code)
	fmt.Println("Message ", response.Result[0].Message)

	//close connection
	client.Conn.Close()
}
