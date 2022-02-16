package frames

/**
created by ahmad ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5732 for more detail
*/
// HostCheckType represents a host check command.
type HostCheckType struct {
	Check HostCheck `xml:"urn:ietf:params:xml:ns:host-1.0 command>check>check"`
}

// HostCheck represents a host check request to the EPP server.
type HostCheck struct {
	Names []string `xml:"name"`
}

func (h *HostCheckType) SetMultiHost(hosts []string) {
	h.Check.Names = hosts
}

// HostInfoType represents a host info command.
type HostInfoType struct {
	Info HostInfo `xml:"urn:ietf:params:xml:ns:host-1.0 command>info>info"`
}

// HostInfo represents a host info request to the EPP server.
type HostInfo struct {
	Name string `xml:"name"`
}

func (h *HostInfoType) SetHost(host string) {
	h.Info.Name = host
}

// HostDeleteType represents a host delete command.
type HostDeleteType struct {
	Delete HostDelete `xml:"urn:ietf:params:xml:ns:host-1.0 command>delete>delete"`
}

// HostDelete represents a host delete request to the EPP server.
type HostDelete struct {
	Name string `xml:"name"`
}

func (h *HostDeleteType) SetHost(host string) {
	h.Delete.Name = host
}
