package frames

/**
Created by ahmad ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5731#section-3.2.1 for more detail
*/
// DomainCreateType implements extension for create from domain-1.0.
type DomainCreateType struct {
	Create    DomainCreate `xml:"urn:ietf:params:xml:ns:domain-1.0 command>create>create"`
	Extension *interface{} `xml:"command>extension"` //this extension for dnssec
}

// DomainCreate represents a domain create command.
type DomainCreate struct {
	Name       string     `xml:"name"`
	Period     Period     `xml:"period,omitempty"`
	NameServer NameServer `xml:"ns,omitempty"`
	Registrant string     `xml:"registrant,omitempty"`
	Contacts   []Contact  `xml:"contact,omitempty"`
	AuthInfo   *AuthInfo  `xml:"authInfo,omitempty"`
}

// Period represents the period unit and value.
type Period struct {
	Value int    `xml:",chardata"`
	Unit  string `xml:"unit,attr"`
}

// NameServer represents a name server for a domain.
type NameServer struct {
	HostObject    []string        `xml:"hostObj,omitempty"`
	HostAttribute []HostAttribute `xml:"hostAttr,omitempty"`
}

// HostAttribute represents attributes for a host for a domain.
type HostAttribute struct {
	HostName    string        `xml:"hostName"`
	HostAddress []HostAddress `xml:"hostAddr"`
}

// Contact represents a contact for a domain.
type Contact struct {
	Name string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

func (d *DomainCreateType) SetDomain(domain string) {
	d.Create.Name = domain
}
func (d *DomainCreateType) SetPeriod(period Period) {
	d.Create.Period = period
}
func (d *DomainCreateType) SetRegistrant(registrant string) {
	d.Create.Registrant = registrant
}
func (d *DomainCreateType) AddContact(contact Contact) {
	d.Create.Contacts = append(d.Create.Contacts, contact)
}
func (d *DomainCreateType) SetAuthInfo(password string) {
	AuthInfo := AuthInfo{}
	AuthInfo.Password = password
	d.Create.AuthInfo = &AuthInfo
}

//make new struct
var NewNameServer NameServer

//add ns without ip address
func (d *DomainCreateType) AddAddr(host string) {
	var NameServer *NameServer = &NewNameServer
	NameServer.HostObject = append(NameServer.HostObject, host)
	d.Create.NameServer = *NameServer
}

//add ns with ip
func (d *DomainCreateType) AddHostObject(host HostAttribute) {
	var NameServer *NameServer = &NewNameServer
	NameServer.HostAttribute = append(NameServer.HostAttribute, host)
	d.Create.NameServer = *NameServer
}

//add dnsecc
func (d *DomainCreateType) AddDnssec(dnssec interface{}) {
	// var dnSEC DNSSECExtensionCreateType

	d.Extension = &dnssec
}
