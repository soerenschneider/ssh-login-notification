package pkg

import (
	"gitlab.com/soerenschneider/ssh-login-notification/internal"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/dispatcher/telegram"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/enrichers/dns"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/enrichers/geo"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/scrapers"
	"testing"
)

func Test_cortex_Run(t *testing.T) {
	options := &internal.Options{}
	options.DnsLookup = true
	options.GeoLookup = true
	cortex := NewCortex(options)

	scraper := &scrapers.DummyIpScraper{}
	cortex.scraper = scraper
	cortex.dnsEnricher = &dns.DnsDummyProvider{}
	cortex.geoEnricher = &geo.GeoDummyEnricher{}

	telegramMock := telegram.MockDispatcher{}
	cortex.dispatcher = &telegramMock
	telegramMock.On("Send", "New login on test for soeren from 1.1.1.1 (reverse, Org, ISP) City, Region, Country").Return(nil)

	cortex.Run()
	telegramMock.AssertExpectations(t)
}
