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

// UserNotification is an interface to dispatch the formatted
// message to the end-user.
type UserNotification interface {
	Send(message string) error
}

// Scraper is an interface that collects the bare information that's
// available on this system about the remote user.
type Scraper interface {
	GetRemoteUserInfo(login *internal.RemoteUserInfo) error
}

// cortex struct configures all the pluggable components that are used
// to perform this workload.
type cortex struct {
	dispatcher  UserNotification
	scraper     Scraper
	geoEnricher GeoEnricher
	dnsEnricher DnsEnricher
	options     *internal.Options
}

// NewCortex initialises this struct.
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

// Run performs the whole workload of this program in an nicely
// abstracted way.
func (c *cortex) Run() {
	login := internal.RemoteUserInfo{}
	login.Host, _ = os.Hostname()

	err := c.scraper.GetRemoteUserInfo(&login)
	if err == nil {
		// If no error occurred while getting remote user info, enrich the
		// information by performing lookups.
		aggregator := NewAggregator(c.options, c.geoEnricher, c.dnsEnricher)
		aggregator.Enrich(&login)
	}

	// format the message and dispatch it
	formatted := formatter.Format(login)
	err = c.dispatcher.Send(formatted)
	if err != nil {
		log.Error("Error while dispatching message")
		os.Exit(1)
	}
}
