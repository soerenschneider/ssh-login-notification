package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sshnot/internal"
	"time"
)

type geoProviderIpInfo struct {
	dnsProvider DnsProvider
}

type DnsProvider interface {
	IpLookup(dns string) (string, error)
}

// NewGeoProviderIpInfo instantiates a new ip geo provider
// that queries ipinfo.io.
func NewGeoProviderIpInfo(dnsProvider *DnsProvider) *geoProviderIpInfo {
	return &geoProviderIpInfo{dnsProvider: *dnsProvider}
}

func (p *geoProviderIpInfo) getIp(remoteHost *RemoteHost) (string, error) {
	if remoteHost.IsIp {
		return remoteHost.Host, nil
	}
	return p.dnsProvider.IpLookup(remoteHost.Host)
}

func (p *geoProviderIpInfo) Lookup(remoteHost *RemoteHost) (*internal.IpGeoInfo, error) {
	timeout := time.Duration(2 * time.Second)
	client := http.Client{Timeout: timeout}

	url := "https://ipinfo.io/%v/geo"

	ip, _ := p.getIp(remoteHost)
	resp, err := client.Get(fmt.Sprintf(url, ip))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := internal.IpGeoInfo{}
	err = json.Unmarshal(body, &ret)

	return &ret, err
}
