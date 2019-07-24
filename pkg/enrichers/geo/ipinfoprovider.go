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
	dnsProvider dnsProvider
}

type dnsProvider interface {
	ResolveIp(dns string) (string, error)
}

// NewGeoProviderIpInfo instantiates a new ip geo provider
// that queries ipinfo.io. Given the facts that we may only
// be able to acquire DNS data for the remote host and that
// ipinfo's API does only work with IP data, we need a DNS
// provider, so that we can resolve the IP address if the
// remote host is indeed only a hostname.
func NewGeoProviderIpInfo(dnsProvider *dnsProvider) *geoProviderIpInfo {
	return &geoProviderIpInfo{dnsProvider: *dnsProvider}
}

func (p *geoProviderIpInfo) getIp(remoteHost *RemoteHost) (string, error) {
	if remoteHost.IsIp {
		return remoteHost.Host, nil
	}
	return p.dnsProvider.ResolveIp(remoteHost.Host)
}

// Lookup performs a lookup on a remote host to gather geo information about the
// appropriate host.
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
