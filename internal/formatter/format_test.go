package formatter

import (
	"fmt"
	"sshnot/internal"
	"testing"
)

func TestFormat(t *testing.T) {
	scrape := internal.SshLoginNotification{}
	scrape.Dns = "dns.tld"
	scrape.Host = "thishost"
	scrape.Ip = "8.8.8.8"
	scrape.User = "soeren"
	scrape.IpInfo.Country = "US"
	scrape.IpInfo.Region = "NJ"
	scrape.IpInfo.City = "Newark"
	scrape.IpInfo.Isp = "Level 3 Communications"
	scrape.IpInfo.Org = "Google Inc"

	actual := Format(scrape)
	expected := "New login on thishost for soeren from 8.8.8.8 (dns.tld, Google Inc, Level 3 Communications)  Newark, NJ, US"
	fmt.Println(actual)
	fmt.Println(expected)

	if expected != actual {
		t.Errorf("Didn't get expected text, got: %v", actual)
	}
}
