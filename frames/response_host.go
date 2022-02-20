package frames

import "time"

/**
response check host
*/
// HostCheckDataType represents host check data.
type HostCheckDataType struct {
	CheckData HostCheckData `xml:"urn:ietf:params:xml:ns:host-1.0 chkData"`
}

// HostCheckData represents the response for a host check command.
type HostCheckData struct {
	Name []CheckType `xml:"cd"`
}

/**
Response info host
*/
// HostInfoDataType represents host info data.
type HostInfoDataType struct {
	InfoData HostInfoData `xml:"urn:ietf:params:xml:ns:host-1.0 infData"`
}

// HostInfoData represents the response for a host info command.
type HostInfoData struct {
	Name         string        `xml:"name"`
	ROID         string        `xml:"roid"`
	Status       []HostStatus  `xml:"status"`
	Address      []HostAddress `xml:"addr,omitempty"`
	ClientID     string        `xml:"clID"`
	CreateID     string        `xml:"crID"`
	CreateDate   time.Time     `xml:"crDate"`
	UpdateID     string        `xml:"upID,omitempty"`
	UpdateDate   time.Time     `xml:"upDate,omitempty"`
	TransferDate time.Time     `xml:"trDate,omitempty"`
}

/**
response create host
*/
// HostCreateDataType represents host create data.
type HostCreateDataType struct {
	CreateData HostCreateData `xml:"urn:ietf:params:xml:ns:host-1.0 creData"`
}

// HostCreateData represents the response for a host create command.
type HostCreateData struct {
	Name       string    `xml:"name"`
	CreateDate time.Time `xml:"crDate"`
}
