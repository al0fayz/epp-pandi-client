package examples

import (
	"encoding/xml"
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
	//print xml
	fmt.Println(string(resCheck))
	response := frames.Response{
		ResultData: &frames.HostCheckDataType{},
	}
	if err := xml.Unmarshal(resCheck, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code is ", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)
	checkData := response.ResultData.(*frames.HostCheckDataType).CheckData.Name
	for _, host := range checkData {
		fmt.Println(host.Name.Value)
		fmt.Println(host.Name.Available)
	}
	client.Conn.Close() //close connection
}
