package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func CreateHost() {
	host := "ns4.bejak.id"
	hostFrame := frames.HostCreateType{}
	hostFrame.SetHost(host)
	addr1 := frames.HostAddress{
		Address: "45.10.10.100",
		IP:      frames.HostIPv4,
	}
	hostFrame.AddAddr(addr1)
	addr2 := frames.HostAddress{
		Address: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		IP:      frames.HostIPv6,
	}
	hostFrame.AddAddr(addr2)
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//create
	resCreate, err := client.CreateHost(&hostFrame)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resCreate))

	client.Conn.Close() //close connection
}