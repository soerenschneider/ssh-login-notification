package formatter

import (
	"fmt"
	"gitlab.com/soerenschneider/ssh-login-notification/internal"
	"testing"
)

func TestFormatFull(t *testing.T) {
	scrape := internal.RemoteUserInfo{}
	scrape.Dns = "dns.tld"
	scrape.Host = "thishost"
	scrape.Ip = "8.8.8.8"
	scrape.User = "soeren"

	scrape.Geo.Country = "US"
	scrape.Geo.Region = "NJ"
	scrape.Geo.City = "Newark"
	scrape.Geo.Isp = "Level 3 Communications"
	scrape.Geo.Org = "Google Inc"

	actual := Format(scrape)
	expected := fmt.Sprintf("New login on %v for %v from %v (%v, %v, %v) %v, %v, %v", scrape.Host, scrape.User, scrape.Ip, scrape.Dns, scrape.Geo.Org, scrape.Geo.Isp, scrape.Geo.City, scrape.Geo.Region, scrape.Geo.Country)

	if expected != actual {
		t.Errorf("Didn't get expected text, got: %v", actual)
	}
}

func TestFormatGeo(t *testing.T) {
	scrape := internal.RemoteUserInfo{}
	scrape.Host = "thishost"
	scrape.Ip = "8.8.8.8"
	scrape.User = "soeren"
	scrape.Geo.Country = "US"
	scrape.Geo.Region = "NJ"
	scrape.Geo.City = "Newark"

	actual := Format(scrape)
	expected := fmt.Sprintf("New login on %v for %v from %v %v, %v, %v", scrape.Host, scrape.User, scrape.Ip, scrape.Geo.City, scrape.Geo.Region, scrape.Geo.Country)

	if expected != actual {
		t.Errorf("Didn't get expected text, got: %v", actual)
	}
}

func TestFormatNoDnsNoIsp(t *testing.T) {
	scrape := internal.RemoteUserInfo{}
	scrape.Host = "thishost"
	scrape.Ip = "8.8.8.8"
	scrape.User = "soeren"
	scrape.Geo.Country = "US"
	scrape.Geo.Region = "NJ"
	scrape.Geo.City = "Newark"
	scrape.Geo.Org = "Google Inc"

	actual := Format(scrape)
	expected := fmt.Sprintf("New login on %v for %v from %v (%v) %v, %v, %v", scrape.Host, scrape.User, scrape.Ip, scrape.Geo.Org, scrape.Geo.City, scrape.Geo.Region, scrape.Geo.Country)

	if expected != actual {
		t.Errorf("Didn't get expected text, got: %v", actual)
	}
}

func TestFormatNoLocation(t *testing.T) {
	scrape := internal.RemoteUserInfo{}
	scrape.Host = "thishost"
	scrape.Ip = "8.8.8.8"
	scrape.User = "soeren"
	scrape.Geo.Isp = "Some ISP"

	actual := Format(scrape)
	expected := fmt.Sprintf("New login on %v for %v from %v (%v)", scrape.Host, scrape.User, scrape.Ip, scrape.Geo.Isp)

	if expected != actual {
		t.Errorf("Didn't get expected text, got: '%v'", actual)
	}
}

func TestFormatBare(t *testing.T) {
	scrape := internal.RemoteUserInfo{}
	scrape.Host = "thishost"
	scrape.Ip = "8.8.8.8"
	scrape.User = "soeren"

	actual := Format(scrape)
	expected := fmt.Sprintf("New login on %v for %v from %v", scrape.Host, scrape.User, scrape.Ip)

	if expected != actual {
		t.Errorf("Didn't get expected text, got: '%v'", actual)
	}
}
