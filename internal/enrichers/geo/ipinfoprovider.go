package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type geoProviderIpInfo struct {
}

func NewGeoProviderIpInfo() *geoProviderIpInfo {
	return &geoProviderIpInfo{}
}

func (p *geoProviderIpInfo) Lookup(ip string) (*IpGeoInfo, error) {
	uri := fmt.Sprintf("https://ipinfo.io/%v/geo", ip)
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
		log.Infof("Received reply: '%v'", string(body))
		return &ret, nil
	}

	return nil, err
}
