package frames

type Logout struct {
	Logout  ClTRID `xml:"command>logout>clTRID,omitempty"`
}
type ClTRID struct {
	Value string `xml:"clTRID,omitempty"`
}
