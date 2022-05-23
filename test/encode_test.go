package test

import (
	"testing"

	"github.com/al0fayz/epp-pandi-client/frames"

	"github.com/al0fayz/epp-pandi-client/epp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	dc := frames.DomainCreateType{
		Create: frames.DomainCreate{
			Name: "example.net",
			Period: frames.Period{
				Value: 12,
				Unit:  "m",
			},
			NameServer: frames.NameServer{
				HostObject: []string{
					"ns1.example.net",
					"ns2.example.net",
				},
			},
			Registrant: "registrant-00001",
			Contacts: []frames.Contact{
				{
					Name: "contact-00001",
					Type: "tech",
				},
				{
					Name: "contact-00002",
					Type: "admin",
				},
			},
			AuthInfo: &frames.AuthInfo{
				Password: "some-password",
			},
		},
	}

	expectedXMLWithNS := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	<command>
		<create>
		<domain:create xmlns:domain="urn:ietf:params:xml:ns:domain-1.0" xmlns="urn:ietf:params:xml:ns:domain-1.0">
			<domain:name>example.net</domain:name>
			<domain:period unit="m">12</domain:period>
			<domain:ns>
			<domain:hostObj>ns1.example.net</domain:hostObj>
			<domain:hostObj>ns2.example.net</domain:hostObj>
			</domain:ns>
			<domain:registrant>registrant-00001</domain:registrant>
			<domain:contact type="tech">contact-00001</domain:contact>
			<domain:contact type="admin">contact-00002</domain:contact>
			<domain:authInfo>
			<domain:pw>some-password</domain:pw>
			</domain:authInfo>
		</domain:create>
		</create>
	</command>
	</epp>
`)

	encoded, err := epp.Encode(dc, epp.ClientXMLAttributes())

	require.Nil(t, err)
	assert.Equal(t, string(expectedXMLWithNS), string(encoded))
}
