package internal

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sshnot/internal/enrichers"
	"sshnot/internal/enrichers/geo"
	"strings"
	"time"
)

const (
	undef = "UNDEF"
)

type Scrape struct {
	Login   *SshLoginNotification
	Options *Options
}

type GeoProvider interface {
	Lookup(ip string) (*geo.IpGeoInfo, error)
}

func NewScrape(options *Options) *Scrape {
	s := Scrape{Login: &SshLoginNotification{}}

	s.Login.Host, _ = os.Hostname()
	s.Login.Ip = undef
	s.Login.User = undef

	s.Login.Date = time.Now()
	s.Options = options

	s.scrapeEnvInfo()

	return &s
}

func (s *Scrape) scrapeEnvInfo() {
	scrapeIp(s)

	if s.Login.Ip == undef && s.Login.Dns == undef {
		return
	}

	ipGeoInfoChan := make(chan *geo.IpGeoInfo)
	dnsChan := make(chan string)

	if s.Options.GeoLookup {
		go fetchIpInfo(s.Login.Ip, ipGeoInfoChan)
	}

	if s.Options.DnsLookup && s.Login.Dns == undef {
		go fetchDns(s.Login.Ip, dnsChan)
	}

	if s.Options.GeoLookup {
		ipGeoInfo := <-ipGeoInfoChan
		s.Login.Geo = *ipGeoInfo
	}

	if s.Options.DnsLookup && s.Login.Dns == undef {
		s.Login.Dns, _ = <-dnsChan
	}

	close(ipGeoInfoChan)
	close(dnsChan)
}

func scrapeIp(scrape *Scrape) {
	extractSuccessful := extractSshClient(scrape)

	if !extractSuccessful {
		extractSuccessful = extractPamRhost(scrape)
	}

	if !extractSuccessful {
		log.Warnf("No info found in SSH_CLIENT and PAM_RHOST")
	}
}

func extractSshClient(scrape *Scrape) bool {
	sshClient := os.Getenv("SSH_CLIENT")
	if len(sshClient) > 0 {
		split := strings.Split(sshClient, " ")
		if len(split) >= 2 {
			scrape.Login.Ip = split[0]
			scrape.Login.Port = split[1]
			scrape.Login.User = os.Getenv("USER")

			return true
		}
	}

	return false
}

func extractPamRhost(scrape *Scrape) bool {
	rhost := os.Getenv("PAM_RHOST")
	if len(rhost) > 0 {
		scrape.Login.Dns = rhost
		scrape.Login.User = os.Getenv("PAM_USER")
		scrape.Login.Ip, _ = enrichers.IpLookup(rhost)

		return true
	}

	return false
}

func fetchIpInfo(s string, c chan *geo.IpGeoInfo) {
	var lookupProvider GeoProvider
	lookupProvider = geo.NewProviderIpApi()
	a, _ := lookupProvider.Lookup(s)
	c <- a
}

func fetchDns(s string, c chan string) {
	dns, err := enrichers.DnsLookup(s)
	if err == nil {
		c <- dns
	}
}
