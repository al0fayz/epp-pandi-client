package frames

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
}

func (h *HostUpdateType) SetHost(host string) {
	h.Update.Name = host
}
func (h *HostUpdateType) AddAddr(host HostAddress) {
	h.Update.Add.Address = append(h.Update.Add.Address, host)
}
func (h *HostUpdateType) RemoveAddr(host HostAddress) {
	h.Update.Remove.Address = append(h.Update.Remove.Address, host)
}
