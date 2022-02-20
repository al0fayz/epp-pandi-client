package frames

import "time"

/**
Response Check Contact
*/
// ContactCheckDataType represents contact check data.
type ContactCheckDataType struct {
	CheckData ContactCheckData `xml:"urn:ietf:params:xml:ns:contact-1.0 chkData"`
}

// ContactCheckData represents the data returned from a contact check command.
type ContactCheckData struct {
	Name []CheckContact `xml:"cd"`
}

// CheckContact represents the data from a contact check command name.
type CheckContact struct {
	Name   CheckName `xml:"id"`
	Reason string    `xml:"reason,omitempty"`
}

/**
response create contact
*/
// ContactCreateDataType represents contact create data.
type ContactCreateDataType struct {
	CreateData ContactCreateData `xml:"urn:ietf:params:xml:ns:contact-1.0 creData"`
}

// ContactCreateData represents the data returned from a contact create command.
type ContactCreateData struct {
	Name       string    `xml:"id"`
	CreateDate time.Time `xml:"crDate"`
}

/**
response info contact
*/
// ContactInfoDataType represents  contact info data.
type ContactInfoDataType struct {
	InfoData ContactInfoData `xml:"urn:ietf:params:xml:ns:contact-1.0 infData"`
}

// ContactInfoData represents the data returned from a contact info command.
type ContactInfoData struct {
	Name         string          `xml:"id"`
	ROID         string          `xml:"roid"`
	Status       []ContactStatus `xml:"status"`
	PostalInfo   []PostalInfo    `xml:"postalInfo"`
	Voice        E164Type        `xml:"voice,omitempty"`
	Fax          E164Type        `xml:"fax,omitempty"`
	Email        string          `xml:"email"`
	ClientID     string          `xml:"clID"`
	CreateID     string          `xml:"crID"`
	CreateDate   time.Time       `xml:"crDate"`
	UpdateID     string          `xml:"upID,omitempty"`
	UpdateDate   time.Time       `xml:"upDate,omitempty"`
	TransferDate time.Time       `xml:"trDate,omitempty"`
	AuthInfo     AuthInfo        `xml:"authInfo,omitempty"`
	Disclose     Disclose        `xml:"disclose,omitempty"`
}
