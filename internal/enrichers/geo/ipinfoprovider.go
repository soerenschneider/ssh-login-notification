package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type geoProviderIpInfo struct {
}

// NewGeoProviderIpInfo instantiates a new ip geo provider
// that queries ipinfo.io.
func NewGeoProviderIpInfo() *geoProviderIpInfo {
	return &geoProviderIpInfo{}
}

func (p *geoProviderIpInfo) Lookup(ip string) (*IpGeoInfo, error) {
	timeout := time.Duration(2 * time.Second)
	client := http.Client{Timeout: timeout}

	url := "https://ipinfo.io/%v/geo"
	resp, err := client.Get(fmt.Sprintf(url, ip))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := IpGeoInfo{}
	err = json.Unmarshal(body, &ret)

	return &ret, err
}
