package frames

import "time"

/**
created by ahmad ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5731 for more detail
*/
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

/**
Renew Domain
see https://datatracker.ietf.org/doc/html/rfc5731#section-3.2.3 for more detail
*/
// DomainRenewType implements extension for renew from domain-1.0.
type DomainRenewType struct {
	Renew DomainRenew `xml:"urn:ietf:params:xml:ns:domain-1.0 command>renew>renew"`
}

// DomainRenew represents a domain renew command.
type DomainRenew struct {
	Name       string    `xml:"name"`
	ExpireDate time.Time `xml:"curExpDate"`
	Period     Period    `xml:"period,omitempty"`
}

func (d *DomainRenewType) SetDomain(domain string) {
	d.Renew.Name = domain
}
func (d *DomainRenewType) SetExpDate(exp time.Time) {
	d.Renew.ExpireDate = exp
}
func (d *DomainRenewType) SetPeriod(period Period) {
	d.Renew.Period = period
}

/**
Delete Domain
see https://datatracker.ietf.org/doc/html/rfc5731#section-3.2.2 for more detail
*/
// DomainDeleteType implements extension for delete from domain-1.0.
type DomainDeleteType struct {
	Delete DomainDelete `xml:"urn:ietf:params:xml:ns:domain-1.0 command>delete>delete"`
}

// DomainDelete represents a domain delete command.
type DomainDelete struct {
	Name string `xml:"name"`
}

func (d *DomainDeleteType) SetDomain(domain string) {
	d.Delete.Name = domain
}
