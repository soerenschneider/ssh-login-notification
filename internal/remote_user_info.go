package internal

import (
	"errors"
	"net"
	"strings"
)

var privateIPBlocks []*net.IPNet

// RemoteUserInfo contains the gathered information about a SSH login.
type RemoteUserInfo struct {
	Ip   string
	Host string
	User string
	Dns  string
	Geo  IpGeoInfo
}

func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

// IsPrivateIP returns whether the login comes from a private
// network or not.
func (s *RemoteUserInfo) IsPrivateIP() (bool, error) {
	ip := net.ParseIP(s.Ip)

	if ip == nil {
		return false, errors.New("can't parse IP")
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true, nil
		}
	}
	return false, nil
}

// PrettyPrintLocation returns a comma separated string of the
// geo information.
func (s *RemoteUserInfo) PrettyPrintLocation() string {
	fields := []string{s.Geo.City, s.Geo.Region, s.Geo.Country}
	return getConcatFields(fields)
}

// PrettyPrintProvider returns a comma separated string of the
// provider information.
func (s *RemoteUserInfo) PrettyPrintProvider() string {
	fields := []string{s.Dns, s.Geo.Org, s.Geo.Isp}
	return getConcatFields(fields)
}

func getConcatFields(fields []string) string {
	nonEmtpy := []string{}
	for _, n := range fields {
		if len(n) > 0 {
			nonEmtpy = append(nonEmtpy, n)
		}
	}
	loc := strings.Join(nonEmtpy, ", ")
	return loc
}
