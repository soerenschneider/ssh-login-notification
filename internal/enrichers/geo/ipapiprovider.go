package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type geoProviderIpApi struct {
}

func New() *geoProviderIpApi {
	return &geoProviderIpApi{}
}

func (p *geoProviderIpApi) Lookup(ip string) (*IpGeoInfo, error) {
	uri := fmt.Sprintf("http://ip-api.com/json/%v", ip)
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	ret := IpGeoInfo{}
	err = json.Unmarshal(body, &ret)

	if err == nil {
		log.Debugf("Received reply: '%v'", string(body))
		return &ret, nil
	}

	return nil, err
}
