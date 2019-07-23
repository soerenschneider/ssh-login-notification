package pkg

import (
	"sshnot/internal"
	"sshnot/pkg/dispatcher/telegram"
	"sshnot/pkg/enrichers/dns"
	"sshnot/pkg/enrichers/geo"
	"sshnot/pkg/scrapers"
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
	cortex.geoEnricher = &geo.GeoDummyProvider{}

	telegramMock, _ := telegram.NewTelegramMock(options)
	cortex.dispatcher = telegramMock
	telegramMock.On("Send", "New login on test for soeren from 1.1.1.1 (reverse, Org, ISP) City, Region, Country").Return(nil)

	cortex.Run()
	telegramMock.AssertExpectations(t)
}
