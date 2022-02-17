package main

import (
	"epp-pandi-client/epp"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
		<response>
		<result code="1000">
			<msg lang="en">Command completed successfully</msg>
		</result>
		<trID>
			<clTRID>ABC-12345</clTRID>
			<svTRID>54321-XYZ</svTRID>
		</trID>
		</response>
	</epp>`)
	response, err := epp.Decode(xml)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
}
func TestDecode2(t *testing.T) {
	//xml example with resData
	xml := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd">
	  <response>
		<result code="1000">
		  <msg>Command completed successfully</msg>
		</result>
		<resData>
		  <domain:infData xmlns:domain="urn:ietf:params:xml:ns:domain-1.0" xsi:schemaLocation="urn:ietf:params:xml:ns:domain-1.0 domain-1.0.xsd">
			<domain:name>pandi.id</domain:name>
			<domain:roid>PANDI-DO173387</domain:roid>
			<domain:status s="clientTransferProhibited"/>
			<domain:status s="clientUpdateProhibited"/>
			<domain:status s="clientDeleteProhibited"/>
			<domain:status s="serverDeleteProhibited"/>
			<domain:status s="serverTransferProhibited"/>
			<domain:status s="serverUpdateProhibited"/>
			<domain:contact type="admin">pndi1</domain:contact>
			<domain:contact type="billing">pndi1</domain:contact>
			<domain:contact type="tech">pndi1</domain:contact>
			<domain:registrant>pndi1</domain:registrant>
			<domain:ns>
			  <domain:hostObj>dns1.pandi.id</domain:hostObj>
			  <domain:hostObj>dns2.pandi.id</domain:hostObj>
			</domain:ns>
			<domain:host>ns1.pandi.id</domain:host>
			<domain:clID>pandi</domain:clID>
			<domain:crID>registrar-console</domain:crID>
			<domain:crDate>2013-04-14T07:27:32.0Z</domain:crDate>
			<domain:exDate>2020-04-14T23:59:59.0Z</domain:exDate>
			<domain:upID>H5364127</domain:upID>
			<domain:upDate>2021-12-09T01:27:20.0Z</domain:upDate>
		  </domain:infData>
		</resData>
		<extension>
		  <secDNS:infData xmlns:secDNS="urn:ietf:params:xml:ns:secDNS-1.1"/>
		</extension>
		<trID>
		  <svTRID>864fc9b6-ba9a-4458-a320-f37fce24f9f8</svTRID>
		  <clTRID>null</clTRID>
		</trID>
	  </response>
	</epp>
	`)
	response, err := epp.Decode(xml)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(response.ResultData)
	assert.Equal(t, 1000, response.Result[0].Code)
	assert.Equal(t, "Command completed successfully", response.Result[0].Message)
}
