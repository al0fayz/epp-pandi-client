package frames

// ContactCheckType represents a contact check command.
type ContactCheckType struct {
	Check ContactCheck `xml:"urn:ietf:params:xml:ns:contact-1.0 command>check>check"`
}

// ContactCheck represents a check for contact(s).
type ContactCheck struct {
	Names []string `xml:"id"`
}

func (c *ContactCheckType) SetMultiContact(contacts []string) {
	c.Check.Names = contacts
}

// ContactInfoType represents a contact info command.
type ContactInfoType struct {
	Info ContactInfo `xml:"urn:ietf:params:xml:ns:contact-1.0 command>info>info"`
}

// ContactInfo represents a contact info command.
type ContactInfo struct {
	Name     string   `xml:"id"`
	AuthInfo AuthInfo `xml:"authInfo,omitempty"`
}

// AuthInfo represents authentication information used when transferring domain.
type AuthInfo struct {
	Password  string `xml:"pw,omitempty"`
	Extension string `xml:"ext,omitempty"`
}

func (c *ContactInfoType) SetContact(contactId string) {
	c.Info.Name = contactId
}
func (c *ContactInfoType) SetAuthInfo(password string) {
	c.Info.AuthInfo.Password = password
}
