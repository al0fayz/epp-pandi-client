package examples

import (
	"epp-pandi-client/epp"
	"epp-pandi-client/frames"
)

func UpdateHost() {
	hostFrame := frames.HostUpdateType{}
	hostFrame.SetHost("ns1.bejak.id")
	addAddr := frames.HostAddress{
		Address: "127.0.0.1",
		IP:      frames.HostIPv4,
	}
	hostFrame.AddAddr(addAddr)
	epp.UpdateHost(&hostFrame)
}
