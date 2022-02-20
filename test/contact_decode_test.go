package test

import (
	"encoding/xml"
	"epp-pandi-client/frames"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactCheck(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <contact:chkData
	       xmlns:contact="urn:ietf:params:xml:ns:contact-1.0">
	        <contact:cd>
	          <contact:id avail="1">sh8013</contact:id>
	        </contact:cd>
	        <contact:cd>
	          <contact:id avail="0">sah8013</contact:id>
	          <contact:reason>In use</contact:reason>
	        </contact:cd>
	        <contact:cd>
	          <contact:id avail="1">8013sah</contact:id>
	        </contact:cd>
	      </contact:chkData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)

	response := frames.Response{
		ResultData: &frames.ContactCheckDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	// json, err := json.MarshalIndent(response, "", "  ")
	// if err != nil {
	// 	t.Error(err)
	// }
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	checkData := response.ResultData.(*frames.ContactCheckDataType).CheckData
	// fmt.Println(checkData.Name[0].Name.Value)
	assert.Equal(t, "sh8013", checkData.Name[0].Name.Value)
	assert.Equal(t, true, checkData.Name[0].Name.Available)
	//contact 2
	assert.Equal(t, "sah8013", checkData.Name[1].Name.Value)
	assert.Equal(t, false, checkData.Name[1].Name.Available)
	assert.Equal(t, "In use", checkData.Name[1].Reason)
	//contact 3
	assert.Equal(t, "8013sah", checkData.Name[2].Name.Value)
	assert.Equal(t, true, checkData.Name[2].Name.Available)

}
func TestContactInfo(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <contact:infData
	       xmlns:contact="urn:ietf:params:xml:ns:contact-1.0">
	        <contact:id>sh8013</contact:id>
	        <contact:roid>SH8013-REP</contact:roid>
	        <contact:status s="linked"/>
	        <contact:status s="clientDeleteProhibited"/>
	        <contact:postalInfo type="int">
	          <contact:name>John Doe</contact:name>
             <contact:org>Example Inc.</contact:org>
             <contact:addr>
               <contact:street>123 Example Dr.</contact:street>
               <contact:street>Suite 100</contact:street>
               <contact:city>Dulles</contact:city>
               <contact:sp>VA</contact:sp>
               <contact:pc>20166-6503</contact:pc>
               <contact:cc>US</contact:cc>
             </contact:addr>
           </contact:postalInfo>
           <contact:voice x="1234">+1.7035555555</contact:voice>
           <contact:fax>+1.7035555556</contact:fax>
           <contact:email>jdoe@example.com</contact:email>
           <contact:clID>ClientY</contact:clID>
           <contact:crID>ClientX</contact:crID>
           <contact:crDate>1999-04-03T22:00:00.0Z</contact:crDate>
           <contact:upID>ClientX</contact:upID>
           <contact:upDate>1999-12-03T09:00:00.0Z</contact:upDate>
           <contact:trDate>2000-04-08T09:00:00.0Z</contact:trDate>
           <contact:authInfo>
             <contact:pw>2fooBAR</contact:pw>
           </contact:authInfo>
           <contact:disclose flag="0">
             <contact:voice/>
             <contact:email/>
           </contact:disclose>
         </contact:infData>
       </resData>
       <trID>
         <clTRID>ABC-12345</clTRID>
         <svTRID>54322-XYZ</svTRID>
       </trID>
     </response>
   </epp>`)
	response := frames.Response{
		ResultData: &frames.ContactInfoDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	infoData := response.ResultData.(*frames.ContactInfoDataType).InfoData
	assert.Equal(t, "sh8013", infoData.Name)
	assert.Equal(t, "SH8013-REP", infoData.ROID)
	assert.Equal(t, frames.ContactStatusLinked, infoData.Status[0].ContactStatusType)
	assert.Equal(t, frames.ContactStatusClientDeleteProhibited, infoData.Status[1].ContactStatusType)
	//postal info
	assert.Equal(t, "John Doe", infoData.PostalInfo[0].Name)
	assert.Equal(t, "Example Inc.", infoData.PostalInfo[0].Organization)
	//address
	assert.Equal(t, "Dulles", infoData.PostalInfo[0].Address.City)
	assert.Equal(t, "US", infoData.PostalInfo[0].Address.CountryCode)
	assert.Equal(t, "20166-6503", infoData.PostalInfo[0].Address.PostalCode)
	assert.Equal(t, "VA", infoData.PostalInfo[0].Address.StateProvince)

	assert.Equal(t, "+1.7035555555", infoData.Voice.Value)
	assert.Equal(t, "+1.7035555556", infoData.Fax.Value)
	assert.Equal(t, "jdoe@example.com", infoData.Email)
	assert.Equal(t, "2fooBAR", infoData.AuthInfo.Password)

}
func TestContactTransfer(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <contact:trnData xmlns:contact="urn:ietf:params:xml:ns:contact-1.0">
	        <contact:id>sh8013</contact:id>
	        <contact:trStatus>pending</contact:trStatus>
	        <contact:reID>ClientX</contact:reID>
	        <contact:reDate>2000-06-06T22:00:00.0Z</contact:reDate>
	        <contact:acID>ClientY</contact:acID>
	        <contact:acDate>2000-06-11T22:00:00.0Z</contact:acDate>
	      </contact:trnData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{
		ResultData: &frames.ContactTransferDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	transferData := response.ResultData.(*frames.ContactTransferDataType).TransferData
	assert.Equal(t, "sh8013", transferData.Name)
	assert.Equal(t, "ClientX", transferData.RequestingID)
	assert.Equal(t, "ClientY", transferData.ActingID)
	assert.Equal(t, frames.ContactTransferPending, transferData.TransferStatus)

}
func TestContactCreate(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <contact:creData
	       xmlns:contact="urn:ietf:params:xml:ns:contact-1.0">
           <contact:id>sh8013</contact:id>
           <contact:crDate>1999-04-03T22:00:00.0Z</contact:crDate>
         </contact:creData>
       </resData>
       <trID>
         <clTRID>ABC-12345</clTRID>
         <svTRID>54321-XYZ</svTRID>
       </trID>
     </response>
   </epp>`)
	response := frames.Response{
		ResultData: &frames.ContactCreateDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	createData := response.ResultData.(*frames.ContactCreateDataType).CreateData
	assert.Equal(t, "sh8013", createData.Name)

}
func TestContactDelete(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54321-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
}
func TestContactUpdate(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
         <msg>Command completed successfully</msg>
       </result>
       <trID>
         <clTRID>ABC-12345</clTRID>
         <svTRID>54321-XYZ</svTRID>
       </trID>
     </response>
   </epp>`)
	response := frames.Response{}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
}
