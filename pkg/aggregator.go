package pkg

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/soerenschneider/ssh-login-notification/internal"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/enrichers/geo"
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
	ResolveIp(dns string) (string, error)
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
func (agg *Aggregator) Enrich(login *internal.RemoteUserInfo) {
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

	if agg.options.GeoLookup {
		go agg.fetchIpInfo(login, ipGeoInfoChan)
	}

	if agg.options.DnsLookup && login.Dns == "" {
		go agg.fetchDns(login.Ip, dnsChan)
	}

	if agg.options.GeoLookup {
		ipGeoInfo := <-ipGeoInfoChan
		if ipGeoInfo != nil {
			login.Geo = *ipGeoInfo

			if login.Ip == "" && login.Geo.Ip != "" {
				login.Ip = login.Geo.Ip
			}
		}
	}

	if agg.options.DnsLookup && login.Dns == "" {
		login.Dns, _ = <-dnsChan
	}
}

func (agg *Aggregator) fetchIpInfo(login *internal.RemoteUserInfo, channel chan *internal.IpGeoInfo) {
	host := geo.RemoteHost{}
	if login.Ip != "" {
		host.IsIp = true
		host.Host = login.Ip
	} else {
		host.Host = login.Dns
	}

	geo, _ := agg.geoEnricher.Lookup(&host)
	channel <- geo
}

func (agg *Aggregator) fetchDns(ip string, channel chan string) {
	dns, _ := agg.dnsEnricher.DnsLookup(ip)
	channel <- dns
}
