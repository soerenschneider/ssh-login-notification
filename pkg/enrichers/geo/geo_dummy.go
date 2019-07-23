package geo

import "sshnot/internal"

type GeoDummyProvider struct {
}

func (p *GeoDummyProvider) Lookup(host *RemoteHost) (*internal.IpGeoInfo, error) {
	info := internal.IpGeoInfo{}

	info.Isp = "ISP"
	info.Org = "Org"
	info.City = "City"
	info.Region = "Region"
	info.Country = "Country"
	info.Ip = host.Host

	return &info, nil
}
