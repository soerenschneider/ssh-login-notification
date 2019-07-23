package pkg

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sshnot/internal"
	"sshnot/pkg/dispatcher/telegram"
	"sshnot/pkg/enrichers/dns"
	"sshnot/pkg/enrichers/geo"
	"sshnot/pkg/formatter"
	"sshnot/pkg/scrapers"
)

type Bot interface {
	Send(message string) error
}

type Scraper interface {
	GetRemoteUserInfo(login *internal.RemoteUserInfo) error
}

type cortex struct {
	dispatcher  Bot
	scraper     Scraper
	geoEnricher GeoEnricher
	dnsEnricher DnsEnricher
	options     *internal.Options
}

func NewCortex(options *internal.Options) *cortex {
	dispatcher, _ := telegram.NewTelegramBot(options)

	return &cortex{
		dispatcher:  dispatcher,
		scraper:     &scrapers.EnvScraper{},
		geoEnricher: geo.NewProviderIpApi(),
		dnsEnricher: &dns.DnsProvider{},
		options:     options,
	}
}

// Run ties all the components together and performs the
// whole workload.
func (c *cortex) Run() {
	login := internal.RemoteUserInfo{}

	err := c.scraper.GetRemoteUserInfo(&login)
	if err == nil {
		// If no error occurred while getting remote user info, enrich the
		// information by performing lookups.
		aggregator := NewAggregator(c.options, c.geoEnricher, c.dnsEnricher)
		aggregator.Enrich(&login)
	}

	formatted := formatter.Format(login)
	err = c.dispatcher.Send(formatted)
	if err != nil {
		log.Error("Error while dispatching message")
		os.Exit(1)
	}
}
