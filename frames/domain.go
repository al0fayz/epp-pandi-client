package frames

// DomainCheckType implements extension for check from domain-1.0.
type DomainCheckType struct {
	Check DomainCheck `xml:"urn:ietf:params:xml:ns:domain-1.0 command>check>check"`
}

// DomainCheck represents a check for domain(s).
type DomainCheck struct {
	Names []string `xml:"name"`
}

func (d *DomainCheckType) SetMultiDomain(domains []string) {
	d.Check.Names = domains
}

// DomainInfoType implements extension for info from domain-1.0.
type DomainInfoType struct {
	Info DomainInfo `xml:"urn:ietf:params:xml:ns:domain-1.0 command>info>info"`
}

// DomainInfo represents a domain info command.
type DomainInfo struct {
	Name     DomainInfoName `xml:"name"`
	AuthInfo *AuthInfo      `xml:"authInfo,omitempty"`
}

// DomainInfoName represents a domain name in a domain info response.
type DomainInfoName struct {
	Name  string          `xml:",chardata"`
	Hosts DomainHostsType `xml:"hosts,attr,omitempty"`
}

// DomainHostsType represents the string value of host types.
type DomainHostsType string

// Constants representing the string value of host value types.
const (
	DomainHostsAll  DomainHostsType = "all"
	DomainHostsDel  DomainHostsType = "del"
	DomainHostsNone DomainHostsType = "none"
	DomainHostsSub  DomainHostsType = "sub"
)

func (d *DomainInfoType) SetDomain(domain string) {
	d.Info.Name.Name = domain
}
func (d *DomainInfoType) SetAuthInfo(password string) {
	d.Info.AuthInfo.Password = password
}
