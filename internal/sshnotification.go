package internal

import (
	"sshnot/internal/enrichers/geo"
	"strings"
	"time"
)

type SshLoginNotification struct {
	Ip   string
	Host string
	User string
	Date time.Time
	Dns  string
	Geo  geo.IpGeoInfo
}

func (s *SshLoginNotification) PrettyPrintLocation() string {
	fields := []string{s.Geo.City, s.Geo.Region, s.Geo.Country}
	return getConcatFields(fields)
}

func (s *SshLoginNotification) PrettyPrintProvider() string {
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
