package frames

/**
created by ahmad ropai <akhmadrofai@gmail.com>
see https://datatracker.ietf.org/doc/html/rfc5733#section-3.2.1 for more detail
*/
// ContactCreateType represents a contact create command.
type ContactCreateType struct {
	Create ContactCreate `xml:"urn:ietf:params:xml:ns:contact-1.0 command>create>create"`
}

// ContactCreate represents a contact create command.
type ContactCreate struct {
	ID         string       `xml:"id"`
	PostalInfo []PostalInfo `xml:"postalInfo"`
	Voice      E164Type     `xml:"voice,omitempty"`
	Fax        E164Type     `xml:"fax,omitempty"`
	Email      string       `xml:"email"`
	AuthInfo   AuthInfo     `xml:"authInfo"`
	Disclose   Disclose     `xml:"disclose,omitempty"`
}

// PostalInfo represents potal information for a contact.
type PostalInfo struct {
	Name         string         `xml:"name"`
	Organization string         `xml:"org,omitempty"`
	Address      Address        `xml:"addr"`
	Type         PostalInfoType `xml:"type,attr"`
}

// Disclose represents fields that may be disclosed to the public.
type Disclose struct {
	Name         InternationalOrLocalType `xml:"name,omitempty"`
	Organization InternationalOrLocalType `xml:"org,omitempty"`
	Address      InternationalOrLocalType `xml:"addr,omitempty"`
	Voice        bool                     `xml:"voice,omitempty"`
	Fax          bool                     `xml:"fax,omitempty"`
	Email        bool                     `xml:"email,omitempty"`
	Flag         bool                     `xml:"flag,attr"`
}

// InternationalOrLocalType represents a value with a type set to an available
// postal info type.
type InternationalOrLocalType struct {
	Value string         `xml:",chardata"`
	Type  PostalInfoType `xml:"type,attr"`
}

// PostalInfoType represents the typoe of a postal info.
type PostalInfoType string

// Contants represeting postal info types.
const (
	PostalInfoLocal         PostalInfoType = "loc"
	PostalInfoInternational PostalInfoType = "int"
)

// Address represents a full address for postal information.
type Address struct {
	Street        []string `xml:"street,omitempty"`
	City          string   `xml:"city"`
	StateProvince string   `xml:"sp,omitempty"`
	PostalCode    string   `xml:"pc,omitempty"`
	CountryCode   string   `xml:"cc"`
}

// E164Type represents the E.164 numeric plan.
//see https://datatracker.ietf.org/doc/html/rfc5733#section-3.2.1
type E164Type struct {
	Value string `xml:",chardata"`
	X     string `xml:"x,attr"`
}

//create
func (c *ContactCreateType) SetContactId(id string) {
	c.Create.ID = id
}

func (c *ContactCreateType) SetAuthInfo(password string) {
	c.Create.AuthInfo.Password = password
}
func (c *ContactCreateType) SetEmail(email string) {
	c.Create.Email = email
}
func (c *ContactCreateType) AddVoice(voice E164Type) {
	c.Create.Voice = voice
}
func (c *ContactCreateType) AddFax(fax E164Type) {
	c.Create.Fax = fax
}

//add postal info
func (c *ContactCreateType) AddPostalInfo(postal PostalInfo) {
	c.Create.PostalInfo = append(c.Create.PostalInfo, postal)
}

//add disclose
func (c *ContactCreateType) AddDisclose(disclose Disclose) {
	c.Create.Disclose = disclose
}
