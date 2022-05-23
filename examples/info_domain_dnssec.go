package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"
)

func InfoDomainDnssec() {
	domain := "bejo5.my.id"
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
	//print xml
	fmt.Println(string(resInfo))

	response := frames.Response{
		ResultData: &frames.DomainInfoDataType{},
		Extension:  &frames.DNSSECExtensionInfoDataType{},
	}
	if err := xml.Unmarshal(resInfo, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Code", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)
	infoData := response.ResultData.(*frames.DomainInfoDataType).InfoData
	fmt.Println("Domain", infoData.Name)
	fmt.Println(infoData.ROID)
	//status
	status := infoData.Status
	for _, st := range status {
		// fmt.Println(st.Status)
		fmt.Println(st.DomainStatusType)
	}
	//contact
	contact := infoData.Contact
	for _, ct := range contact {
		fmt.Println(ct.Name)
		fmt.Println(ct.Type)
	}
	//registrant
	fmt.Println(infoData.Registrant)
	//expire
	fmt.Println(infoData.ExpireDate)
	//client id
	fmt.Println(infoData.ClientID)
	//host
	for _, host := range infoData.Host {
		fmt.Println(host)
	}
	//auth
	fmt.Println(infoData.AuthInfo.Password)
	//nameserver
	for _, ns := range infoData.NameServer.HostAttribute {
		fmt.Println("ns", ns.HostAddress)
		fmt.Println("ns", ns.HostName)
	}
	//hostObjt
	for _, ho := range infoData.NameServer.HostObject {
		fmt.Println("host object", ho)
	}

	//dnssec
	ext := response.Extension.(*frames.DNSSECExtensionInfoDataType).InfoData
	fmt.Println("MaxSignatureLife :", ext.MaxSignatureLife)
	for _, ds := range ext.DNSSECData {
		fmt.Println(ds.KeyTag)
		fmt.Println(ds.Algorithm)
		fmt.Println(ds.DigestType)
		fmt.Println(ds.Digest)
	}
	client.Conn.Close() //close connection
}
