package frames

const (
	NameSpaceRGP10 = "urn:ietf:params:xml:ns:rgp-1.0"
)

// restore domain
type DomainRestoreType struct {
	Update    DomainRestore  `xml:"urn:ietf:params:xml:ns:domain-1.0 command>update>update"`
	Extension *RgpUpdateType `xml:"command>extension"`
}
type DomainRestore struct {
	Name string `xml:"command>update>update>name"`
}
type RgpUpdateType struct {
	Rgp RgpType `xml:"urn:ietf:params:xml:ns:rgp-1.0 update"`
}
type RgpType struct {
	Restore Restore `xml:"restore"`
}
type Restore struct {
	Op string `xml:"op,attr"`
}

func (d *DomainRestoreType) SetDomain(domain string) {
	d.Update.Name = domain
}
