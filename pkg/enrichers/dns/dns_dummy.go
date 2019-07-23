package dns

type DnsDummyProvider struct {
}

// DnsLookup accepts an ip address and performs a reverse dns lookup.
func (this *DnsDummyProvider) DnsLookup(ip string) (string, error) {
	return "reverse", nil
}

// IpLookup accepts a hostname and performs a dns lookup to resolve its ip
// address.
func (this *DnsDummyProvider) IpLookup(dns string) (string, error) {
	return "1.1.1.1", nil
}
