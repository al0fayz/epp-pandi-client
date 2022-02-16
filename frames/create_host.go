package frames

/**
Created by ahmad ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5732#section-3.2.1 for more detail
*/
// HostCreateType represents a host create command.
type HostCreateType struct {
	Create HostCreate `xml:"urn:ietf:params:xml:ns:host-1.0 command>create>create"`
}

// HostCreate represents a host create request to the EPP server.
type HostCreate struct {
	Name    string        `xml:"name"`
	Address []HostAddress `xml:"addr,omitempty"`
}

// HostAddress represents an IP address beloning to a host.
type HostAddress struct {
	Address string `xml:",chardata"`
	IP      IPType `xml:"ip,attr"`
}

// IPType represents IP type, that is the IP version
type IPType string

// Contants representing allowed values for IP type.
const (
	HostIPv4 IPType = "v4"
	HostIPv6 IPType = "v6"
)

func (h *HostCreateType) SetHost(host string) {
	h.Create.Name = host
}

func (h *HostCreateType) AddAddr(host HostAddress) {
	h.Create.Address = append(h.Create.Address, host)
}
