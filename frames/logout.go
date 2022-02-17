package frames

type Logout struct {
	ClientID string `xml:"command>logout>clID,omitempty"`
}
