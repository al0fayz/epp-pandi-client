package frames

/**
Created by ahmad ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5731#section-3.2.5 for more detail

*/
// DomainStatusType represents available status values.
type DomainStatusType string

// Constants representing the string value of status value types.
const (
	DomainStatusClientDeleteProhibited   DomainStatusType = "clientDeleteProhibited"
	DomainStatusClientHold               DomainStatusType = "clientHold"
	DomainStatusClientRenewProhibited    DomainStatusType = "clientRenewProhibited"
	DomainStatusClientTransferProhibited DomainStatusType = "clientTransferProhibited"
	DomainStatusClientUpdateProhibited   DomainStatusType = "clientUpdateProhibited"
	DomainStatusInactive                 DomainStatusType = "inactive"
	DomainStatusOk                       DomainStatusType = "ok"
	DomainStatusPendingCreate            DomainStatusType = "pendingCreate"
	DomainStatusPendingDelete            DomainStatusType = "pendingDelete"
	DomainStatusPendingRenew             DomainStatusType = "pendingRenew"
	DomainStatusPendingTransfer          DomainStatusType = "pendingTransfer"
	DomainStatusPendingUpdate            DomainStatusType = "pendingUpdate"
	DomainStatusServerDeleteProhibited   DomainStatusType = "serverDeleteProhibited"
	DomainStatusServerHold               DomainStatusType = "serverHold"
	DomainStatusServerRenewProhibited    DomainStatusType = "serverRenewProhibited"
	DomainStatusServerTransferProhibited DomainStatusType = "serverTransferProhibited"
	DomainStatusServerUpdateProhibited   DomainStatusType = "serverUpdateProhibited"
)

// DomainUpdateType implements extension for update from domain-1.0.
type DomainUpdateType struct {
	Update DomainUpdate `xml:"urn:ietf:params:xml:ns:domain-1.0 command>update>update"`
}

// DomainUpdate represents a domain update command.
type DomainUpdate struct {
	Name   string           `xml:"command>update>update>name"`
	Add    *DomainAddRemove `xml:"command>update>update>add,omitempty"`
	Remove *DomainAddRemove `xml:"command>update>update>rem,omitempty"`
	Change *DomainChange    `xml:"command>update>update>chg,omitempty"`
}

// DomainAddRemove ...
type DomainAddRemove struct {
	NameServer NameServer     `xml:"ns,omitempty"`
	Contact    []Contact      `xml:"contact,omitempty"`
	Status     []DomainStatus `xml:"status,omitempty"`
}

// DomainChange ...
type DomainChange struct {
	Registrant string    `xml:"registrant,omitempty"`
	AuthInfo   *AuthInfo `xml:"authInfo,omitempty"`
}

// DomainStatus represents statuses for a domain.
type DomainStatus struct {
	Status           string           `xml:",chardata"`
	DomainStatusType DomainStatusType `xml:"s,attr"`
	Language         string           `xml:"lang,attr,omitempty"`
}

func (d *DomainUpdateType) SetDomain(domain string) {
	d.Update.Name = domain
}

//change data
func (d *DomainUpdateType) ChangeData(data DomainChange) {
	var domainChange *DomainChange = &data
	d.Update.Change = domainChange
}

//remove data
func (d *DomainUpdateType) RemoveData(data DomainAddRemove) {
	var domainAddremove *DomainAddRemove = &data
	d.Update.Remove = domainAddremove
}

//add data
func (d *DomainUpdateType) AddData(data DomainAddRemove) {
	var domainAddremove *DomainAddRemove = &data
	d.Update.Add = domainAddremove
}
