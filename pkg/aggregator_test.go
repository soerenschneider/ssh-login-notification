package pkg

import (
	"os"
	"sshnot/internal"
	"sshnot/pkg/scrapers"
	"testing"
)

func TestScrapeIpV6(t *testing.T) {
	raw := "::1 55234 22"
	os.Setenv("SSH_CLIENT", raw)

	scraper := scrapers.EnvScraper{}
	remoteUserInfo := internal.RemoteUserInfo{}
	err := scraper.GetRemoteUserInfo(&remoteUserInfo)

	if err != nil {
		t.Error("Error occured")
	}

	if remoteUserInfo.Ip != "::1" {
		t.Errorf("Expected Ip to be ::1 but is: '%v'", remoteUserInfo.Ip)
	}

	os.Unsetenv("SSH_CLIENT")
}

func TestScrapeIpV4(t *testing.T) {
	raw := "123.123.123.123 55234 22"
	os.Setenv("SSH_CLIENT", raw)

	scraper := scrapers.EnvScraper{}
	remoteUserInfo := internal.RemoteUserInfo{}
	err := scraper.GetRemoteUserInfo(&remoteUserInfo)

	if err != nil {
		t.Error("Error occured")
	}

	if remoteUserInfo.Ip != "123.123.123.123" {
		t.Errorf("Expected Ip to be 123.123.123.123 but is: '%v'", remoteUserInfo.Ip)
	}

	os.Unsetenv("SSH_CLIENT")
}

func TestScrapeRhost(t *testing.T) {
	raw := "localhost"
	os.Setenv("PAM_RHOST", raw)

	scraper := scrapers.EnvScraper{}
	remoteUserInfo := internal.RemoteUserInfo{}
	err := scraper.GetRemoteUserInfo(&remoteUserInfo)

	if err != nil {
		t.Error("Error occured")
	}

	if remoteUserInfo.Ip != "127.0.0.1" && remoteUserInfo.Ip != "::1" {
		t.Errorf("Expected Ip to be 123.123.123.123 but is: '%v'", remoteUserInfo.Ip)
	}

	os.Unsetenv("PAM_RHOST")
}
