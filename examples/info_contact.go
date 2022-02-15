package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func InfoContact() {
	contactId := "epp-598"
	contact := frames.ContactInfoType{}
	contact.SetContact(contactId)
	// contact.SetAuthInfo("password")
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//check contact
	resInfo, err := client.InfoContact(&contact)
	if err != nil {
		fmt.Println(err)
	}
	//response check contact
	fmt.Println(string(resInfo))

	client.Conn.Close() //close connection
}
