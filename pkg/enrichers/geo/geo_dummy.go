package geo

import "sshnot/internal"

// GeoDummyEnricher is a dummy that delivers fixed data for an arbitrary host.
type GeoDummyEnricher struct {
}

// Lookup performs a lookup on a remote host to gather geo information about the
// appropriate host.
func (enricher *GeoDummyEnricher) Lookup(host *RemoteHost) (*internal.IpGeoInfo, error) {
	info := internal.IpGeoInfo{}

	info.Isp = "ISP"
	info.Org = "Org"
	info.City = "City"
	info.Region = "Region"
	info.Country = "Country"
	info.Ip = host.Host

	return &info, nil
}
