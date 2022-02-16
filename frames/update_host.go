package frames

/**
Created by ahmad ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5732#section-3.2.5 for more detail
*/
// HostUpdateType represents a host update command.
type HostUpdateType struct {
	Update HostUpdate `xml:"urn:ietf:params:xml:ns:host-1.0 command>update>update"`
}

// HostUpdate represents a host update request to the EPP server.
type HostUpdate struct {
	Name   string         `xml:"name"`
	Add    *HostAddRemove `xml:"add,omitempty"`
	Remove *HostAddRemove `xml:"rem,omitempty"`
	Change string         `xml:"chg>name,omitempty"`
}

// HostAddRemove represents data that can be added or removed while updating a
// domain.
type HostAddRemove struct {
	Address []HostAddress `xml:"addr,omitempty"`
	Status  []HostStatus  `xml:"status,omitempty"`
}

// HostStatus represents statuses for a host.
type HostStatus struct {
	Status         string         `xml:",chardata"`
	HostStatusType HostStatusType `xml:"s,attr"`
	Language       string         `xml:"lang,attr,omitempty"`
}

// HostStatusType represents available status values.
type HostStatusType string

// Constants representing the string value of status value types.
const (
	HostStatusClientDeleteProhibited HostStatusType = "clientDeleteProhibited"
	HostStatusClientUpdateProhibited HostStatusType = "clientUpdateProhibited"
	HostStatusLinked                 HostStatusType = "linked"
	HostStatusOk                     HostStatusType = "ok"
	HostStatusPendingCreate          HostStatusType = "pendingCreate"
	HostStatusPendingDelete          HostStatusType = "pendingDelete"
	HostStatusPendingTransfer        HostStatusType = "pendingTransfer"
	HostStatusPendingUpdate          HostStatusType = "pendingUpdate"
	HostStatusServerDeleteProhibited HostStatusType = "serverDeleteProhibited"
	HostStatusServerUpdateProhibited HostStatusType = "serverUpdateProhibited"
)

//make global variabel for make blue print of HostAddRemove
var AddHostAddRemove HostAddRemove    //for add
var RemoveHostAddRemove HostAddRemove //for remove

func (h *HostUpdateType) SetHost(host string) {
	h.Update.Name = host
}
func (h *HostUpdateType) AddAddr(host HostAddress) {
	//copy blue print AddHostAddRemove
	var HostAddRemove *HostAddRemove = &AddHostAddRemove
	HostAddRemove.Address = append(HostAddRemove.Address, host)
	h.Update.Add = HostAddRemove
}
func (h *HostUpdateType) RemoveAddr(host HostAddress) {
	//copy blue print RemoveHostAddRemove
	var HostAddRemove *HostAddRemove = &RemoveHostAddRemove
	HostAddRemove.Address = append(HostAddRemove.Address, host)
	h.Update.Remove = HostAddRemove
}
func (h *HostUpdateType) ChangeHost(host string) {
	h.Update.Change = host
}

//add status host
func (h *HostUpdateType) AddStatus(status HostStatus) {
	//copy blue print AddHostAddRemove
	var HostAddRemove *HostAddRemove = &AddHostAddRemove
	HostAddRemove.Status = append(HostAddRemove.Status, status)
	h.Update.Add = HostAddRemove
}

//remove status host
func (h *HostUpdateType) RemoveStatus(status HostStatus) {
	//copy blue print RemoveHostAddRemove
	var HostAddRemove *HostAddRemove = &RemoveHostAddRemove
	HostAddRemove.Status = append(HostAddRemove.Status, status)
	h.Update.Remove = HostAddRemove
}
