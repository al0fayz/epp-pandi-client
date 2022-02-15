package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func CheckHost() {
	hosts := []string{
		"ns1.pandi.id",
		"ns1.ayam.id",
	}
	hostFrame := frames.HostCheckType{}
	hostFrame.SetMultiHost(hosts)
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//check host
	resCheck, err := client.CheckHost(&hostFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response check host
	fmt.Println(string(resCheck))
	client.Conn.Close() //close connection
}
