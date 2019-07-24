package dns

// DnsDummyProvider is a mock that fulfills the DNS enricher interface.
type DnsDummyProvider struct {
}

// DnsLookup accepts an ip address and performs a reverse dns lookup.
func (this *DnsDummyProvider) DnsLookup(ip string) (string, error) {
	return "reverse", nil
}

// ResolveIp accepts a hostname and performs a dns lookup to resolve its ip
// address.
func (this *DnsDummyProvider) ResolveIp(dns string) (string, error) {
	return "1.1.1.1", nil
}
