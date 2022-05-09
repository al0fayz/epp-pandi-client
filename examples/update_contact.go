package examples

import (
	"encoding/xml"
	"epp-pandi-client/epp"
	"epp-pandi-client/frames"
	"fmt"
)

func UpdateContact() {
	contactFrame := frames.ContactUpdateType{}
	contactFrame.SetContactId("hello123")
	//change data
	//change voice
	voice := frames.E164Type{
		Value: "087804650415",
	}
	contactFrame.ChangeVoice(voice)
	//add fax
	fax := frames.E164Type{
		Value: "087804650415",
	}
	contactFrame.ChangeFax(fax)
	//add address
	postalInfo := frames.PostalInfo{
		Name:         "alfa code",
		Organization: "pandi",
		Type:         frames.PostalInfoInternational, //this default value == int
		// Type: frames.PostalInfoLocal, //this local == loc
		Address: frames.Address{
			//required
			City:        "indonesia",
			CountryCode: "ID",
			//optional
			StateProvince: "Jawa Barat",
			PostalCode:    "123456",
			Street: []string{
				"jln satu langkah",
				"rt 01 rw 02",
			},
		},
	}
	contactFrame.ChangePostalInfo(postalInfo)

	contactFrame.ChangeAuthInfo("pandi123")
	contactFrame.ChangeEmail("alfa@me.com")

	//OPTIONAL
	//add disclose (optional)
	disclose := frames.Disclose{
		Email: false,
		Voice: true,
	}
	contactFrame.ChangeDisclose(disclose)
	//end change data

	//add status
	status := frames.ContactStatus{}
	status.Status = "ok"
	status.ContactStatusType = frames.ContactStatusOk
	contactFrame.AddStatus(status)

	//remove status
	removeStatus := frames.ContactStatus{}
	removeStatus.Status = "clientDeleteProhibited"
	removeStatus.ContactStatusType = frames.ContactStatusClientDeleteProhibited
	contactFrame.RemoveStatus(removeStatus)
	//END OPTIONAL

	//print xml command update contact
	encoded, err := epp.Encode(contactFrame, epp.ClientXMLAttributes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encoded))

	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//update contact
	resUpdate, err := client.UpdateContact(&contactFrame)
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

	//print
	fmt.Println("code is", response.Result[0].Code)
	fmt.Println("Message ", response.Result[0].Message)

	//close connection
	client.Conn.Close()
}
