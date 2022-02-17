package examples

import (
	"epp-pandi-client/epp"
	"epp-pandi-client/frames"
)

func UpdateDomain() {
	domainFrame := frames.DomainUpdateType{}
	domainFrame.SetDomain("bejak.id")
	epp.UpdateDomain(&domainFrame)

	epp.Logout("hello")
}
