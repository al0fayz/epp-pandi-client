package examples

import (
	"encoding/xml"
	"fmt"

	"github.com/al0fayz/epp-pandi-client/frames"
)

func InfoContact() {
	contactId := "hello123s"
	contact := frames.ContactInfoType{}
	contact.SetContact(contactId)
	// contact.SetAuthInfo("password")
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//check contact
	resInfo, err := client.InfoContact(&contact)
	if err != nil {
		fmt.Println(err)
	}
	//response check contact
	//print xml
	// fmt.Println(string(resInfo))

	response := frames.Response{
		ResultData: &frames.ContactInfoDataType{},
	}
	if err := xml.Unmarshal(resInfo, &response); err != nil {
		fmt.Println(err)
	}
	//print
	fmt.Println("code is", response.Result[0].Code)
	fmt.Println("Message ", response.Result[0].Message)
	//contact info data
	infoData := response.ResultData.(*frames.ContactInfoDataType).InfoData
	fmt.Println(infoData.Name)
	fmt.Println(infoData.ROID) //contact id
	fmt.Println(infoData.Email)
	fmt.Println(infoData.ClientID)
	fmt.Println(infoData.CreateID)
	//status
	fmt.Println(infoData.Status[0].ContactStatusType)
	//authinfo
	fmt.Println(infoData.AuthInfo.Password)
	//phone and fax
	fmt.Println(infoData.Fax.Value)
	fmt.Println(infoData.Voice.Value)
	//address
	fmt.Println(infoData.PostalInfo[0].Address)
	fmt.Println(infoData.PostalInfo[0].Name)
	fmt.Println(infoData.PostalInfo[0].Organization)

	client.Conn.Close() //close connection
}
