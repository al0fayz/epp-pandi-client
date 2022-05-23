package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"
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
	//print xml
	// fmt.Println(string(resInfo))

	response := frames.Response{
		ResultData: &frames.HostInfoDataType{},
	}
	if err := xml.Unmarshal(resInfo, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code is ", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)
	infoData := response.ResultData.(*frames.HostInfoDataType).InfoData
	fmt.Println(infoData.Name)
	fmt.Println(infoData.ClientID)
	fmt.Println(infoData.CreateID)
	fmt.Println(infoData.CreateDate)
	//status
	fmt.Println(infoData.Status[0].Status)
	//addr
	fmt.Println(infoData.Address[0].Address)
	fmt.Println(infoData.Address[0].IP)

	//dnssec
	client.Conn.Close() //close connection
}
