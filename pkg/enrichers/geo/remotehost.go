package geo

// RemoteHost contains the remote host and further information
// whether the host portion is a valid IP address or a hostname.
type RemoteHost struct {
	// Host is the remote host that logged in.
	Host string

	// IsIp denotes whether the host is either a valid IPv4/IPv6
	// address or a hostname.
	IsIp bool
}
