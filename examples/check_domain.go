package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func CheckDomain() {
	domains := []string{
		"pandi.id",
		"ayam-goreng.id",
	}
	domainFrame := frames.DomainCheckType{}
	domainFrame.SetMultiDomain(domains)
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//check domain
	resCheck, err := client.CheckDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response check domain
	fmt.Println(string(resCheck))

	client.Conn.Close() //close connection
}
