package pkg

import (
	log "github.com/sirupsen/logrus"
	"sshnot/internal"
	"sshnot/pkg/enrichers/geo"
)

// Aggregator accepts information about the connecting remote host and
// enriches that data with pluggable providers.
type Aggregator struct {
	geoEnricher GeoEnricher
	dnsEnricher DnsEnricher
	options     *internal.Options
}

// GeoEnricher is used to perform lookups on the host in order
// to get geo information.
type GeoEnricher interface {
	Lookup(host *geo.RemoteHost) (*internal.IpGeoInfo, error)
}

// DnsEnricher is used to perform dns lookups on the host in order
// to get dns information about the host.
type DnsEnricher interface {
	DnsLookup(ip string) (string, error)
	IpLookup(dns string) (string, error)
}

// NewAggregator instantiates a new struct and scrapes the providers
// to collect information about the ip.
func NewAggregator(options *internal.Options, geoProvider GeoEnricher, dnsProvider DnsEnricher) *Aggregator {
	this := Aggregator{
		geoEnricher: geoProvider,
		dnsEnricher: dnsProvider,
		options:     options,
	}

	return &this
}

// Enrich accepts the remote user information and enriches it using the
// configured providers.
func (s *Aggregator) Enrich(login *internal.RemoteUserInfo) {
	if nil == login {
		return
	}

	if login.Ip == "" && login.Dns == "" {
		log.Panic("Everything empty")
	}

	ipGeoInfoChan := make(chan *internal.IpGeoInfo)
	dnsChan := make(chan string)
	defer close(ipGeoInfoChan)
	defer close(dnsChan)

	if s.options.GeoLookup {
		go s.fetchIpInfo(login, ipGeoInfoChan)
	}

	if s.options.DnsLookup && login.Dns == "" {
		go s.fetchDns(login.Ip, dnsChan)
	}

	if s.options.GeoLookup {
		ipGeoInfo := <-ipGeoInfoChan
		if ipGeoInfo != nil {
			login.Geo = *ipGeoInfo

			if login.Ip == "" && login.Geo.Ip != "" {
				login.Ip = login.Geo.Ip
			}
		}
	}

	if s.options.DnsLookup && login.Dns == "" {
		login.Dns, _ = <-dnsChan
	}
}

func (this *Aggregator) fetchIpInfo(login *internal.RemoteUserInfo, channel chan *internal.IpGeoInfo) {
	host := geo.RemoteHost{}
	if login.Ip != "" {
		host.IsIp = true
		host.Host = login.Ip
	} else {
		host.Host = login.Dns
	}

	geo, _ := this.geoEnricher.Lookup(&host)
	channel <- geo
}

func (this *Aggregator) fetchDns(ip string, channel chan string) {
	dns, _ := this.dnsEnricher.DnsLookup(ip)
	channel <- dns
}
