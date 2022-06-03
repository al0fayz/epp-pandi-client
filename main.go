package main

import "github.com/al0fayz/epp-pandi-client/examples"

func main() {
	// examples.ExampleLogin()
	/**
	-> Contact
	*/
	// examples.CheckContact()
	// examples.InfoContact()
	// examples.CreateContact()
	// examples.UpdateContact()
	// examples.DeleteContact()
	// examples.TransferContact()
	/**
	-> Domain
	*/
	// examples.CheckDomain()
	// examples.InfoDomain()
	// examples.InfoDomainDnssec()
	// examples.CreateDomain()
	// examples.CreateDomainDnssec()
	// examples.UpdateDomain()
	// examples.UpdateDomainDnssec()
	// examples.RenewDomain()
	// examples.DeleteDomain()
	// examples.TransferDomain()
	// examples.TransferDomainQuery()
	/**
	-> Host
	*/
	// examples.CheckHost()
	// examples.InfoHost()
	// examples.CreateHost()
	// examples.UpdateHost()
	// examples.DeleteHost()

	//poll
	// examples.PollRequest()
	// examples.PollAck()

	//delete domain
	examples.DeleteDomainByFile()
}
