package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"
)

func UpdateHost() {
	hostFrame := frames.HostUpdateType{}
	hostFrame.SetHost("ns4.bejak.id")
	//add address
	addAddr := frames.HostAddress{
		Address: "127.0.0.1",
		IP:      frames.HostIPv4,
	}
	hostFrame.AddAddr(addAddr)
	//remove address
	removeAddr := frames.HostAddress{
		Address: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		IP:      frames.HostIPv6,
	}
	hostFrame.RemoveAddr(removeAddr)
	//add status
	addStatus := frames.HostStatus{
		Status:         "clientUpdateProhibited",
		HostStatusType: frames.HostStatusClientUpdateProhibited,
	}
	hostFrame.AddStatus(addStatus)
	//remove status
	removeStatus := frames.HostStatus{
		Status:         "pendingCreate",
		HostStatusType: frames.HostStatusPendingDelete,
	}
	hostFrame.RemoveStatus(removeStatus)

	hostFrame.ChangeHost("ns1.bejak.id")
	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//update
	resUpdate, err := client.UpdateHost(&hostFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response update
	//print xml
	fmt.Println(string(resUpdate))
	response := frames.Response{}
	if err := xml.Unmarshal(resUpdate, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code is ", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)

	//close connection
	client.Conn.Close()
}
