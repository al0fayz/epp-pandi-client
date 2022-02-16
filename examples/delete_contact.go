package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func DeleteContact() {
	contactFrame := frames.ContactDeleteType{}
	contactFrame.SetContactId("hello123")
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
	fmt.Println(string(resDelete))

	//close connection
	client.Conn.Close()
}
