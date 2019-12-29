package scrapers

import "gitlab.com/soerenschneider/ssh-login-notification/internal"

// DummyIpScraper is a dummy that provides static data.
type DummyIpScraper struct {
}

// GetRemoteUserInfo scrapes all available information about the remote host and writes
// it into supplied login object.
func (scraper *DummyIpScraper) GetRemoteUserInfo(login *internal.RemoteUserInfo) error {
	login.Ip = "1.1.1.1"
	login.Host = "test"
	login.User = "soeren"

	return nil
}
