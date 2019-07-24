package scrapers

import (
	"errors"
	"net"
	"os"
	"sshnot/internal"
	"strings"
)

// EnvScraper collects information about the remote host from
// all available environment variables.
type EnvScraper struct {
}

func (this *EnvScraper) GetRemoteUserInfo(login *internal.RemoteUserInfo) error {
	extractSuccessful := this.trySshClient(login)

	if !extractSuccessful {
		extractSuccessful = this.tryPam(login)
	}

	var err error = nil
	if !extractSuccessful {
		err = errors.New("No info found in SSH_CLIENT and PAM_RHOST")
	}

	return err
}

// trySshClient collects information from the 'SSH_CLIENT' env variable.
func (this *EnvScraper) trySshClient(login *internal.RemoteUserInfo) bool {
	sshClient := os.Getenv("SSH_CLIENT")
	if len(sshClient) > 0 {
		split := strings.Split(sshClient, " ")
		if len(split) >= 1 {
			login.Ip = split[0]
			login.User = os.Getenv("USER")

			return true
		}
	}

	return false
}

// tryPam collects information from the 'PAM_USER' and 'PAM_RHOST' variables.
func (this *EnvScraper) tryPam(login *internal.RemoteUserInfo) bool {
	if !this.isSessionOpened() {
		return false
	}

	login.User = os.Getenv("PAM_USER")
	rhost := os.Getenv("PAM_RHOST")
	if len(rhost) == 0 {
		return false
	}

	// On some systems this may be either a hostname or an IP.
	// Try to parse as IP, if it doesn't work it's most likely
	// the reverse dns for the host.
	ip := net.ParseIP(rhost)
	if nil == ip {
		login.Dns = rhost
	} else {
		login.Ip = ip.String()
	}

	return true
}

// isSessionOpened checks whether the correct PAM event has happened
// for out notification script.
func (this *EnvScraper) isSessionOpened() bool {
	event := os.Getenv("PAM_TYPE")
	// We are only interested in the "open_session" event. If we don't
	// distinct this, it's possible that messages are being send on
	// disconnects too.
	return len(event) > 1 && event == "open_session"
}