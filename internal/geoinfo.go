package internal

// IpGeoInfo contains all the information that was gathered about
// a IP.
type IpGeoInfo struct {
	Ip      string `json:"query,omitempty"`
	City    string `json:"city,omitempty"`
	Region  string `json:"region,omitempty"`
	Country string `json:"country,omitempty"`
	Isp     string `json:"isp,omitempty"`
	Org     string `json:"org,omitempty"`
}
