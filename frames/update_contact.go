package frames

/**
Created by Ahmad Ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5733#section-3.2.5 for more detail
*/
// ContactStatusType represents contact status types.
type ContactStatusType string

// Constants representing contact status types.
const (
	ContactStatusClientDeleteProhibited   ContactStatusType = "clientDeleteProhibited"
	ContactStatusClientTransferProhibited ContactStatusType = "clientTransferProhibited"
	ContactStatusClientUpdateProhibited   ContactStatusType = "clientUpdateProhibited"
	ContactStatusLinked                   ContactStatusType = "linked"
	ContactStatusOk                       ContactStatusType = "ok"
	ContactStatusPendingCreate            ContactStatusType = "pendingCreate"
	ContactStatusPendingDelete            ContactStatusType = "pendingDelete"
	ContactStatusPendingTransfer          ContactStatusType = "pendingTransfer"
	ContactStatusPendingUpdate            ContactStatusType = "pendingUpdate"
	ContactStatusServerDeleteProhibited   ContactStatusType = "serverDeleteProhibited"
	ContactStatusServerTransferProhibited ContactStatusType = "serverTransferProhibited"
	ContactStatusServerUpdateProhibited   ContactStatusType = "serverUpdateProhibited"
)

// ContactUpdateType represents a contact update command.
type ContactUpdateType struct {
	Update ContactUpdate `xml:"urn:ietf:params:xml:ns:contact-1.0 command>transfer>transfer"`
}

// ContactUpdate represents a contact update command.
type ContactUpdate struct {
	Name   string            `xml:"id"`
	Add    *ContactAddRemove `xml:"add,omitempty"`
	Remove *ContactAddRemove `xml:"rem,omitempty"`
	Change *ContactChange    `xml:"chg>name,omitempty"`
}

// ContactAddRemove represents the fields that holds data to add or remove for a
// contact.
type ContactAddRemove struct {
	Status []ContactStatus `xml:"status"`
}

// ContactStatus represents statuses for a contact.
type ContactStatus struct {
	Status            string            `xml:",chardata"`
	ContactStatusType ContactStatusType `xml:"s,attr"`
	Language          string            `xml:"lang,attr"`
}

// ContactChange represents the data that may be changed while updating a
// contact.
type ContactChange struct {
	PostalInfo []PostalInfo `xml:"postalInfo,omitempty"`
	Voice      E164Type     `xml:"voice,omitempty"`
	Fax        E164Type     `xml:"fax,omitempty"`
	Email      string       `xml:"email,omitempty"`
	AuthInfo   AuthInfo     `xml:"authInfo,omitempty"`
	Disclose   Disclose     `xml:"disclose,omitempty"`
}
