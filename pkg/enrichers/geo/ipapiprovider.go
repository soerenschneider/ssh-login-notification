package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sshnot/internal"
	"time"
)

type geoProviderIpApi struct {
	url string
}

// NewProviderIpApi instantiates a new ip geo provider that queries
// ip-api.com for information about given IP.
func NewProviderIpApi(url ...string) *geoProviderIpApi {
	endpoint := "http://ip-api.com/json/%v"

	if len(url) > 0 {
		endpoint = url[0]
	}

	return &geoProviderIpApi{url: endpoint}
}

func (p *geoProviderIpApi) Lookup(remoteHost *RemoteHost) (*internal.IpGeoInfo, error) {
	timeout := time.Duration(2 * time.Second)
	client := http.Client{Timeout: timeout}

	resp, err := client.Get(fmt.Sprintf(p.url, remoteHost.Host))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := internal.IpGeoInfo{}
	err = json.Unmarshal(body, &ret)

	return &ret, err
}
