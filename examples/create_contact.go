package examples

import (
	"epp-pandi-client/frames"
	"fmt"
)

func CreateContact() {
	contactFrame := frames.ContactCreateType{}
	contactFrame.SetContactId("hello123")
	//add voice
	voice := frames.E164Type{
		Value: "087804650415",
	}
	contactFrame.AddVoice(voice)
	//add fax
	fax := frames.E164Type{
		Value: "087804650415",
	}
	contactFrame.AddFax(fax)
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
	contactFrame.AddPostalInfo(postalInfo)

	contactFrame.SetAuthInfo("pandi123")
	contactFrame.SetEmail("alfa@me.com")
	//add disclose (optional)
	disclose := frames.Disclose{
		Email: false,
		Voice: true,
	}
	contactFrame.AddDisclose(disclose)
	//login
	client, err := Login()
	if err != nil {
		fmt.Println(err)
	}
	//create contact
	resCreate, err := client.CreateContact(&contactFrame)
	if err != nil {
		fmt.Println(err)
	}
	//response create
	fmt.Println(string(resCreate))

	//close connection
	client.Conn.Close()
}
