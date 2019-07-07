package internal

import (
	"sshnot/internal/enrichers/geo"
	"strings"
	"time"
)

type SshLoginNotification struct {
	Ip     string
	Host   string
	Port   string
	User   string
	Date   time.Time
	Dns    string
	IpInfo geo.IpGeoInfo
}

func (s *SshLoginNotification) PrettyPrintLocation() string {
	fields := []string{s.IpInfo.City, s.IpInfo.Region, s.IpInfo.Country}
	return getConcatFields(fields)
}

func (s *SshLoginNotification) PrettyPrintProvider() string {
	fields := []string{s.Dns, s.IpInfo.Org, s.IpInfo.Isp}
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