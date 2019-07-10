package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type geoProviderIpApi struct {
}

func NewProviderIpApi() *geoProviderIpApi {
	return &geoProviderIpApi{}
}

func (p *geoProviderIpApi) Lookup(ip string) (*IpGeoInfo, error) {
	timeout := time.Duration(2 * time.Second)
	client := http.Client{Timeout: timeout}

	url := "http://ip-api.com/json/%v"
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
