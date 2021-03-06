package pkg

import (
	"os"
	"gitlab.com/soerenschneider/ssh-login-notification/internal"
	"gitlab.com/soerenschneider/ssh-login-notification/pkg/scrapers"
	"testing"
)

func TestScrapeIpV6(t *testing.T) {
	raw := "::1 55234 22"
	os.Setenv("SSH_CLIENT", raw)

	scraper := scrapers.EnvScraper{}
	remoteUserInfo := internal.RemoteUserInfo{}
	err := scraper.GetRemoteUserInfo(&remoteUserInfo)

	if err != nil {
		t.Error("Error occurred")
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
		t.Error("Error occurred")
	}

	if remoteUserInfo.Ip != "123.123.123.123" {
		t.Errorf("Expected Ip to be 123.123.123.123 but is: '%v'", remoteUserInfo.Ip)
	}

	os.Unsetenv("SSH_CLIENT")
}
