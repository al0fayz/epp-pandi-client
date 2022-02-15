package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func InfoHost() {
	host := "ns1.pandi.id"
	hostFrame := frames.HostInfoType{}
	hostFrame.SetHost(host)
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//info host
	resInfo, err := client.InfoHost(&hostFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response check host
	fmt.Println(string(resInfo))

	client.Conn.Close() //close connection
}
