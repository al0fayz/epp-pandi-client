package test

import (
	"encoding/xml"
	"epp-pandi-client/frames"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainCheck(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <domain:chkData
	       xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
	        <domain:cd>
	          <domain:name avail="1">example.com</domain:name>
	        </domain:cd>
	        <domain:cd>
	          <domain:name avail="0">example.net</domain:name>
	          <domain:reason>In use</domain:reason>
	        </domain:cd>
	        <domain:cd>
	          <domain:name avail="1">example.org</domain:name>
	        </domain:cd>
	      </domain:chkData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{
		ResultData: &frames.DomainChekDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	checkData := response.ResultData.(*frames.DomainChekDataType).CheckData
	//host 1
	assert.Equal(t, "example.com", checkData.CheckDomain[0].Name.Value)
	assert.Equal(t, true, checkData.CheckDomain[0].Name.Available)
	//host 2
	assert.Equal(t, "example.net", checkData.CheckDomain[1].Name.Value)
	assert.Equal(t, false, checkData.CheckDomain[1].Name.Available)
	assert.Equal(t, "In use", checkData.CheckDomain[1].Reason)
	//host 3
	assert.Equal(t, "example.org", checkData.CheckDomain[2].Name.Value)
	assert.Equal(t, true, checkData.CheckDomain[2].Name.Available)
}

func TestDomainTransfer(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <domain:trnData
	       xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
	        <domain:name>example.com</domain:name>
	        <domain:trStatus>pending</domain:trStatus>
	        <domain:reID>ClientX</domain:reID>
	        <domain:reDate>2000-06-06T22:00:00.0Z</domain:reDate>
	        <domain:acID>ClientY</domain:acID>
	        <domain:acDate>2000-06-11T22:00:00.0Z</domain:acDate>
	        <domain:exDate>2002-09-08T22:00:00.0Z</domain:exDate>
	      </domain:trnData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{
		ResultData: &frames.DomainTransferDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	transferData := response.ResultData.(*frames.DomainTransferDataType).TransferData
	assert.Equal(t, "example.com", transferData.Name)
	assert.Equal(t, frames.DomainTransferPending, transferData.TransferStatus)
	assert.Equal(t, "ClientX", transferData.RequestingID)
	assert.Equal(t, "ClientY", transferData.ActingID)

}
func TestDomainCreate(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <domain:creData
	       xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
	        <domain:name>example.com</domain:name>
	        <domain:crDate>1999-04-03T22:00:00.0Z</domain:crDate>
	        <domain:exDate>2001-04-03T22:00:00.0Z</domain:exDate>
	      </domain:creData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54321-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{
		ResultData: &frames.DomainCreateDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	createData := response.ResultData.(*frames.DomainCreateDataType).CreateData
	assert.Equal(t, "example.com", createData.Name)
}
func TestDomainDelete(t *testing.T) {
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

func TestDomainUpdate(t *testing.T) {
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

func TestDomainRenew(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <domain:renData
	       xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
	        <domain:name>example.com</domain:name>
	        <domain:exDate>2005-04-03T22:00:00.0Z</domain:exDate>
	      </domain:renData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{
		ResultData: &frames.DomainRenewDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	renewData := response.ResultData.(*frames.DomainRenewDataType).RenewData
	assert.Equal(t, "example.com", renewData.Name)
}

func TestDomainInfo(t *testing.T) {
	res_xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	  <response>
	    <result code="1000">
	      <msg>Command completed successfully</msg>
	    </result>
	    <resData>
	      <domain:infData
	       xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
	        <domain:name>example.com</domain:name>
	        <domain:roid>EXAMPLE1-REP</domain:roid>
	        <domain:status s="ok"/>
	        <domain:registrant>jd1234</domain:registrant>
	        <domain:contact type="admin">sh8013</domain:contact>
	        <domain:contact type="tech">sh8013</domain:contact>
	        <domain:ns>
	          <domain:hostObj>ns1.example.com</domain:hostObj>
	          <domain:hostObj>ns1.example.net</domain:hostObj>
	        </domain:ns>
	        <domain:host>ns1.example.com</domain:host>
	        <domain:host>ns2.example.com</domain:host>
	        <domain:clID>ClientX</domain:clID>
	        <domain:crID>ClientY</domain:crID>
	        <domain:crDate>1999-04-03T22:00:00.0Z</domain:crDate>
	        <domain:upID>ClientX</domain:upID>
	        <domain:upDate>1999-12-03T09:00:00.0Z</domain:upDate>
	        <domain:exDate>2005-04-03T22:00:00.0Z</domain:exDate>
	        <domain:trDate>2000-04-08T09:00:00.0Z</domain:trDate>
	        <domain:authInfo>
	          <domain:pw>2fooBAR</domain:pw>
	        </domain:authInfo>
	      </domain:infData>
	    </resData>
	    <trID>
	      <clTRID>ABC-12345</clTRID>
	      <svTRID>54322-XYZ</svTRID>
	    </trID>
	  </response>
	</epp>`)
	response := frames.Response{
		ResultData: &frames.DomainInfoDataType{},
	}
	err := xml.Unmarshal(res_xml, &response)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
	infoData := response.ResultData.(*frames.DomainInfoDataType).InfoData
	assert.Equal(t, "example.com", infoData.Name)
	assert.Equal(t, "EXAMPLE1-REP", infoData.ROID)
	assert.Equal(t, frames.DomainStatusOk, infoData.Status[0].DomainStatusType)
	assert.Equal(t, "jd1234", infoData.Registrant)
	//contact
	assert.Equal(t, "sh8013", infoData.Contact[0].Name)
	assert.Equal(t, "admin", infoData.Contact[0].Type)
	assert.Equal(t, "sh8013", infoData.Contact[1].Name)
	assert.Equal(t, "tech", infoData.Contact[1].Type)
	//host
	assert.Equal(t, "ns1.example.com", infoData.Host[0])
	assert.Equal(t, "ns2.example.com", infoData.Host[1])

	assert.Equal(t, "2fooBAR", infoData.AuthInfo.Password)

}
