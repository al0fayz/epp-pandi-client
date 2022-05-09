package examples

import (
	"encoding/xml"
	"epp-pandi-client/frames"
	"fmt"
	"time"
)

func RenewDomain() {
	domainFrame := frames.DomainRenewType{}
	domainFrame.SetDomain("bejo.my.id")
	exp, err := time.Parse("2006-01-02", "2024-02-18")
	if err != nil {
		fmt.Println(err)
	}
	domainFrame.SetExpDate(exp)
	//period renew
	period := frames.Period{
		Value: 1,
		Unit:  "y", //m or y
	}
	domainFrame.SetPeriod(period)
	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//renew domain
	resRenew, err := client.RenewDomain(&domainFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response error
	//print xml
	fmt.Println(string(resRenew))
	response := frames.Response{
		ResultData: &frames.DomainRenewDataType{},
	}
	if err := xml.Unmarshal(resRenew, &response); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Code", response.Result[0].Code)
	fmt.Println("Message", response.Result[0].Message)
	//domain
	renewData := response.ResultData.(*frames.DomainRenewDataType).RenewData
	fmt.Println("Domain ", renewData.Name)
	fmt.Println("Expire ", renewData.ExpireDate)

	//close connection
	client.Conn.Close()
}
