package examples

import (
	"epp-pandi-client/frames"
	"fmt"
	"time"
)

func RenewDomain() {
	domainFrame := frames.DomainRenewType{}
	domainFrame.SetDomain("bejo.my.id")
	exp, err := time.Parse("2006-01-02", "2029-02-16")
	if err != nil {
		fmt.Println(err)
	}
	domainFrame.SetExpDate(exp)
	//period renew
	period := frames.Period{
		Value: 6,
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
	fmt.Println(string(resRenew))

	//close connection
	client.Conn.Close()
}
