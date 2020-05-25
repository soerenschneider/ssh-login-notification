package pkg

import (
	log "github.com/sirupsen/logrus"
	"os"
	"gitlab.com/soerenschneider/ssh-login-notification/internal"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/dispatcher/telegram"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/enrichers/dns"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/enrichers/geo"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/formatter"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/scrapers"
)

// UserNotification is an interface to dispatch the formatted
// message to the end-user.
type UserNotification interface {
	Send(message string) error
}

// Scraper is an interface that collects the bare information that's
// available on this system about the remote user. If an error is
// returned, the script will not continue.
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
		dnsEnricher: &dns.DnsEnricher{},
		options:     options,
	}
}

// Run performs the whole workload of this program in an nicely
// abstracted way.
func (c *cortex) Run() {
	login := internal.RemoteUserInfo{}
	login.Host, _ = os.Hostname()

	err := c.scraper.GetRemoteUserInfo(&login)
	if err != nil {
		return
	}

	isPrivateIp, _ := login.IsPrivateIP()
	if isPrivateIp && c.options.IgnorePrivateIps {
		return
	}

	if !isPrivateIp {
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
