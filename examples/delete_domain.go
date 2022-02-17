package examples

import (
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
	fmt.Println(string(resDelete))

	//close connection
	// client.Conn.Close()
}
