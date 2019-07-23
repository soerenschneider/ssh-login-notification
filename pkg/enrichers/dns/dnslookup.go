package dns

import (
	"context"
	"net"
	"time"
)

type DnsProvider struct {
}

// DnsLookup accepts an ip address and performs a reverse dns lookup.
func (this *DnsProvider) DnsLookup(ip string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	// avoid resource leaks
	defer cancel()

	var r net.Resolver
	names, err := r.LookupAddr(ctx, ip)
	if err == nil && len(names) > 0 {
		return names[0], err
	}

	return "", err
}

// IpLookup accepts a hostname and performs a dns lookup to resolve its ip
// address.
func (this *DnsProvider) IpLookup(dns string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	// avoid resource leaks
	defer cancel()

	var r net.Resolver
	ip, err := r.LookupIPAddr(ctx, dns)
	if err == nil {
		return ip[0].String(), err
	}

	return "", err
}
