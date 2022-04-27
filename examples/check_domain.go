package examples

import (
	"encoding/xml"
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
	//print xml
	fmt.Println(string(resCheck))
	response := frames.Response{
		ResultData: &frames.DomainChekDataType{},
	}
	if err := xml.Unmarshal(resCheck, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("code is ", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)
	resData := response.ResultData.(*frames.DomainChekDataType).CheckData.CheckDomain
	for _, dom := range resData {
		fmt.Println("Domain ", dom.Name.Value)
		fmt.Println("Avillable ", dom.Name.Available)
		fmt.Println("Reason ", dom.Reason)
	}
	client.Conn.Close() //close connection
}
