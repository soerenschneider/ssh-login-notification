package internal

import (
	"os"
	"testing"
)

func TestScrapeIpV6(t *testing.T) {
	raw := "::1 55234 22"
	os.Setenv("SSH_CLIENT", raw)
	scrape := NewScrape(&Options{GeoLookup: false, DnsLookup: false})

	if scrape.Login == nil {
		t.Error("Scrape is nil")
	}

	if scrape.Login.Ip != "::1" {
		t.Errorf("Expected Ip to be ::1 but is: '%v'", scrape.Login.Ip)
	}
}

func TestScrapeIpV4(t *testing.T) {
	raw := "123.123.123.123 55234 22"
	os.Setenv("SSH_CLIENT", raw)

	scrape := NewScrape(&Options{GeoLookup: false, DnsLookup: false})

	if scrape.Login == nil {
		t.Error("Scrape is nil")
	}

	if scrape.Login.Ip != "123.123.123.123" {
		t.Errorf("Expected Ip to be 123.123.123.123 but is: '%v'", scrape.Login.Ip)
	}
}

func TestScrapeRhost(t *testing.T) {
	raw := "localhost"
	os.Setenv("PAM_RHOST", raw)

	scrape := NewScrape(&Options{GeoLookup: false, DnsLookup: false})

	if scrape.Login == nil {
		t.Error("Scrape is nil")
	}

	if scrape.Login.Ip != "127.0.0.1" && scrape.Login.Ip != "::1" {
		t.Errorf("Expected Ip to be 123.123.123.123 but is: '%v'", scrape.Login.Ip)
	}
}
