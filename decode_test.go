package main

import (
	"fmt"
	"testing"

	"aqwari.net/xml/xmltree"
)

func TestDecode(t *testing.T) {
	response := []byte(
		`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
			<epp xmlns="urn:ietf:params:xml:ns:epp-1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd">
			<response><result code="1000">
			<msg>Command completed successfully</msg>
			</result><trID><svTRID>20806895-826a-44e0-8491-46b2185b86e2</svTRID><clTRID>null</clTRID></trID>
			</response>
		</epp>`)
	root, err := xmltree.Parse(response)
	if err != nil {
		t.Error(err)
	}
	for _, el := range root.Search("", "response") {
		fmt.Println(el)
	}
}
