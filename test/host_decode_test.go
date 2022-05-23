package test

import (
	"encoding/xml"
	"testing"

	"github.com/al0fayz/epp-pandi-client/frames"

	"github.com/stretchr/testify/assert"
)

func TestHostCheck(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <host:chkData
	       xmlns:host="urn:ietf:params:xml:ns:host-1.0">
	        <host:cd>
	          <host:name avail="1">ns1.example.com</host:name>
	        </host:cd>
	        <host:cd>
	          <host:name avail="0">ns2.example2.com</host:name>
	          <host:reason>In use</host:reason>
	        </host:cd>
	        <host:cd>
	          <host:name avail="1">ns3.example3.com</host:name>
	        </host:cd>
	      </host:chkData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)

	response := frames.Response{
		ResultData: &frames.HostCheckDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	checkData := response.ResultData.(*frames.HostCheckDataType).CheckData
	//host 1
	assert.Equal(t, "ns1.example.com", checkData.Name[0].Name.Value)
	assert.Equal(t, true, checkData.Name[0].Name.Available)
	//host 2
	assert.Equal(t, "ns2.example2.com", checkData.Name[1].Name.Value)
	assert.Equal(t, false, checkData.Name[1].Name.Available)
	assert.Equal(t, "In use", checkData.Name[1].Reason)
	//host 3
	assert.Equal(t, "ns3.example3.com", checkData.Name[2].Name.Value)
	assert.Equal(t, true, checkData.Name[2].Name.Available)
}

func TestHostInfo(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <host:infData
	       xmlns:host="urn:ietf:params:xml:ns:host-1.0">
	        <host:name>ns1.example.com</host:name>
	        <host:roid>NS1_EXAMPLE1-REP</host:roid>
	        <host:status s="linked"/>
	        <host:status s="clientUpdateProhibited"/>
	        <host:addr ip="v4">192.0.2.2</host:addr>
	        <host:addr ip="v4">192.0.2.29</host:addr>
	        <host:addr ip="v6">1080:0:0:0:8:800:200C:417A</host:addr>
	        <host:clID>ClientY</host:clID>
	        <host:crID>ClientX</host:crID>
	        <host:crDate>1999-04-03T22:00:00.0Z</host:crDate>
	        <host:upID>ClientX</host:upID>
           <host:upDate>1999-12-03T09:00:00.0Z</host:upDate>
           <host:trDate>2000-04-08T09:00:00.0Z</host:trDate>
         </host:infData>
       </resData>
       <trID>
         <clTRID>ABC-12345</clTRID>
         <svTRID>54322-XYZ</svTRID>
       </trID>
     </response>
   </epp>`)
	response := frames.Response{
		ResultData: &frames.HostInfoDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	infoData := response.ResultData.(*frames.HostInfoDataType).InfoData
	assert.Equal(t, "ns1.example.com", infoData.Name)
	assert.Equal(t, "NS1_EXAMPLE1-REP", infoData.ROID)
	assert.Equal(t, frames.HostStatusLinked, infoData.Status[0].HostStatusType)
	assert.Equal(t, frames.HostStatusClientUpdateProhibited, infoData.Status[1].HostStatusType)
	assert.Equal(t, "ClientY", infoData.ClientID)
	//addr 1
	assert.Equal(t, "192.0.2.2", infoData.Address[0].Address)
	assert.Equal(t, frames.HostIPv4, infoData.Address[0].IP)
	//addr 2
	assert.Equal(t, "192.0.2.29", infoData.Address[1].Address)
	assert.Equal(t, frames.HostIPv4, infoData.Address[1].IP)
	//addr 3
	assert.Equal(t, "1080:0:0:0:8:800:200C:417A", infoData.Address[2].Address)
	assert.Equal(t, frames.HostIPv6, infoData.Address[2].IP)

}
func TestHostCreate(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <host:creData
	       xmlns:host="urn:ietf:params:xml:ns:host-1.0">
	        <host:name>ns1.example.com</host:name>
	        <host:crDate>1999-04-03T22:00:00.0Z</host:crDate>
	      </host:creData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{
		ResultData: &frames.HostCreateDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	createData := response.ResultData.(*frames.HostCreateDataType).CreateData
	assert.Equal(t, "ns1.example.com", createData.Name)
}
func TestHostDelete(t *testing.T) {
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
func TestHostUpdate(t *testing.T) {
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
