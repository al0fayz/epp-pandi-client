# EPP Client


To generate XML to be used for a client, use the specified type for this.

```go
domainInfo := types.DomainInfoType{
    Info: types.DomainInfo{
        Name: types.DomainInfoName{
            Name:  "example.se",
            Hosts: types.DomainHostsAll,
        },
    },
}

bytes, err := Encode(
    domainInfo,
    ClientXMLAttributes(),
)

if err != nil {
    panic(err)
}
```

The above code will generate the following XML.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
  <command>
    <info>
      <domain:info xmlns:domain="urn:ietf:params:xml:ns:domain-1.0" xmlns="urn:ietf:params:xml:ns:domain-1.0">
        <domain:name hosts="all">example.se</domain:name>
      </domain:info>
    </info>
  </command>
</epp>
```

To unmarshal already created XML no matter the namespace or alias, use the auto
genrated types. The XML listed above could be unmarshaled like this.

```go
request := DomainInfoTypeIn{}

if err := xml.Unmarshal(inData, &request); err != nil {
    panic(err)
}

fmt.Println(request.Info.Name.Name) // Prints `example.se`
```

## References

### XSD files

All XSD schemas can be found att [IANA web
page](https://www.iana.org/assignments/xml-registry/xml-registry.xhtml). XSD
files from this repository linked below.

* [contact-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/contact-1.0.xsd)
* [domain-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/domain-1.0.xsd)
* [epp-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/epp-1.0.xsd)
* [eppcom-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/eppcom-1.0.xsd)
* [host-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/host-1.0.xsd)
* [secDNS-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/secDNS-1.0.xsd)
* [secDNS-1.1.xsd](https://www.iana.org/assignments/xml-registry/schema/secDNS-1.1.xsd)

### EPP RFC

* [RFC 5730 Extensible Provisioning Protocol (EPP)](http://www.rfc-editor.org/rfc/rfc5730.txt)
* [RFC 5731 Extensible Provisioning Protocol (EPP) Domain Name Mapping](http://www.rfc-editor.org/rfc/rfc5731.txt)
* [RFC 5732 Extensible Provisioning Protocol (EPP) Host Mapping](http://www.rfc-editor.org/rfc/rfc5732.txt)
* [RFC 5733 Extensible Provisioning Protocol (EPP) Contact Mapping](http://www.rfc-editor.org/rfc/rfc5733.txt)
* [RFC 5734 Extensible Provisioning Protocol (EPP) Transport over TCP](http://www.rfc-editor.org/rfc/rfc5734.txt)
* [RFC 5910 Domain Name System (DNS) Security Extensions Mapping for the Extensible Provisioning Protocol (EPP)](http://www.rfc-editor.org/rfc/rfc5910.txt)

