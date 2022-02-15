package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func DeleteHost() {
	hostFrame := frames.HostDeleteType{}
	hostFrame.SetHost("ns3.bejak.id")
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	resDelete, err := client.DeleteHost(&hostFrame)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resDelete))

	client.Conn.Close() //close connection
}
