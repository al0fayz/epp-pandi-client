package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"
)

func CheckContact() {
	contacts := []string{
		"hello",
		"ayam-goreng",
	}
	contact := frames.ContactCheckType{}
	contact.SetMultiContact(contacts)
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//check contact
	resCheck, err := client.CheckContact(&contact)
	if err != nil {
		fmt.Println(err)
	}
	//response check contact
	//print xml
	// fmt.Println(string(resCheck))

	response := frames.Response{
		ResultData: &frames.ContactCheckDataType{},
	}
	if err := xml.Unmarshal(resCheck, &response); err != nil {
		fmt.Println(err)
	}
	//print object response
	// fmt.Println(response)
	fmt.Println("code is ", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)
	//result contact check
	contactData := response.ResultData.(*frames.ContactCheckDataType).CheckData.Name
	for _, contact := range contactData {
		fmt.Println("Contact Id", contact.Name.Value)
		fmt.Println("Availlable", contact.Name.Available)
	}
	client.Conn.Close() //close connection
}
