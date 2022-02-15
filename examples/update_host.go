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
	removeAddr := frames.HostAddress{
		Address: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		IP:      frames.HostIPv6,
	}
	hostFrame.RemoveAddr(removeAddr)
	hostFrame.ChangeHost("sn1.bejak.id")
	epp.UpdateHost(&hostFrame)
}
