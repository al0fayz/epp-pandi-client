package examples

import (
	"epp-pandi-client/epp"
	"epp-pandi-client/frames"
	"fmt"
)

func TransferDomain() {
	domainFrame := frames.DomainTransferType{
		Transfer: frames.DomainTransfer{
			Name: "aimarch.id",
			Period: frames.Period{
				Value: 1,
				Unit:  "y",
			},
			Authinfo: &frames.AuthInfo{
				Password: "BnS8ycht",
			},
			Operation: "query",
		},
	}
	//print xml command create domain
	encoded, err := epp.Encode(domainFrame, epp.ClientXMLAttributes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encoded))

	// //login
	// client, err := Login()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //create domain
	// resCreate, err := client.TransferDomain(&domainFrame)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //response create
	// //print xml
	// fmt.Println(string(resCreate))

	// //handle response
	// response := frames.Response{}

	// if err := xml.Unmarshal(resCreate, &response); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("code is", response.Result[0].Code)
	// fmt.Println("Message ", response.Result[0].Message)

	// //close connection
	// client.Conn.Close()
}
