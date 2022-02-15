package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func InfoDomain() {
	domain := "pandi.id"
	domainFrame := frames.DomainInfoType{}
	domainFrame.SetDomain(domain)
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//info domain
	resInfo, err := client.InfoDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response info domain
	fmt.Println(string(resInfo))

	client.Conn.Close() //close connection
}
