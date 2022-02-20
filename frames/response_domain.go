package frames

import "time"

/**
response check domain
*/
// DomainChekDataType implements extension for chekData from domain-1.0.
type DomainChekDataType struct {
	CheckData DomainCheckData `xml:"urn:ietf:params:xml:ns:domain-1.0 chkData"`
}

// DomainCheckData represents the response data for a domain check command.
type DomainCheckData struct {
	CheckDomain []CheckType `xml:"cd"`
}

/**
Response create domain
*/
// DomainCreateDataType implements extension for createData from domain-1.0.
type DomainCreateDataType struct {
	CreateData DomainCreateData `xml:"urn:ietf:params:xml:ns:domain-1.0 creData"`
}

// DomainCreateData represents the response data for a domain create command.
type DomainCreateData struct {
	Name       string    `xml:"name"`
	CreateDate time.Time `xml:"crDate"`
	ExpireDate time.Time `xml:"exDate"`
}

/**
response transfer domain
*/
// DomainTransferStatusType represents available transaction statuses.
type DomainTransferStatusType string

// Constants representing the string value of transaction status types.
const (
	DomainTransferClientApproved  DomainTransferStatusType = "clientApproved"
	DomainTransferClientCancelled DomainTransferStatusType = "clientCancelled"
	DomainTransferClientRejected  DomainTransferStatusType = "clientRejected"
	DomainTransferPending         DomainTransferStatusType = "pending"
	DomainTransferServerApproved  DomainTransferStatusType = "serverApproved"
	DomainTransferServerCancelled DomainTransferStatusType = "serverCancelled"
)

// DomainTransferDataType implements extension for transferData from domain-1.0.
type DomainTransferDataType struct {
	TransferData DomainTransferData `xml:"urn:ietf:params:xml:ns:domain-1.0 trnData"`
}

// DomainTransferData represents the response data for a domain transfer command.
type DomainTransferData struct {
	Name           string                   `xml:"name"`
	TransferStatus DomainTransferStatusType `xml:"trStatus"`
	RequestingID   string                   `xml:"reID"`
	RequestingDate string                   `xml:"reDate"`
	ActingID       string                   `xml:"acID"`
	ActingDate     string                   `xml:"acDate"`
	ExpireDate     string                   `xml:"exDate,omitempty"`
}

/**
response renew domain
*/
// DomainRenewDataType implements extension for renewData from domain-1.0.
type DomainRenewDataType struct {
	RenewData DomainRenewData `xml:"urn:ietf:params:xml:ns:domain-1.0 renData"`
}

// DomainRenewData represents the response data for a domain renew command.
type DomainRenewData struct {
	Name       string    `xml:"name"`
	ExpireDate time.Time `xml:"exDate"`
}

/**
response info domain
*/
// DomainInfoDataType implements extension for infoData from domain-1.0.
type DomainInfoDataType struct {
	InfoData DomainInfoData `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
}

// DomainInfoData represents the response data for a domain info command.
type DomainInfoData struct {
	Name         string         `xml:"name"`
	ROID         string         `xml:"roid"`
	Status       []DomainStatus `xml:"status,omitempty"`
	Registrant   string         `xml:"registrant,omitempty"`
	Contact      []Contact      `xml:"contact,omitempty"`
	NameServer   *NameServer    `xml:"ns,omitempty"`
	Host         []string       `xml:"host,omitempty"`
	ClientID     string         `xml:"clID"`
	CreateID     string         `xml:"crID,omitempty"`
	CreateDate   *time.Time     `xml:"crDate,omitempty"`
	UpdateID     string         `xml:"upID,omitempty"`
	UpdateDate   *time.Time     `xml:"upDate,omitempty"`
	ExpireDate   *time.Time     `xml:"exDate,omitempty"`
	TransferDate *time.Time     `xml:"trDate,omitempty"`
	AuthInfo     *AuthInfo      `xml:"authInfo,omitempty"`
}
