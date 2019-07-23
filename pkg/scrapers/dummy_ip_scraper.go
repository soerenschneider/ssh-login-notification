package scrapers

import "sshnot/internal"

type DummyIpScraper struct {
}

// readFromEnv reads information to start with from the environment variables
func (this *DummyIpScraper) GetRemoteUserInfo(login *internal.RemoteUserInfo) error {
	login.Ip = "1.1.1.1"
	login.Host = "test"
	login.User = "soeren"

	return nil
}
