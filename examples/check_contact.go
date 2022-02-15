package examples

import (
	"epp-pandi-client/frames"
	"fmt"
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
	fmt.Println(string(resCheck))
	client.Conn.Close()
}
