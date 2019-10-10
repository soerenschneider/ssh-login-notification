// +build integration

package geo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_geoProviderIpApi_Lookup_Empty(t *testing.T) {
	url := "http://localhost:8080/json/%v"

	provider := NewProviderIpApi(url)

	_, err := provider.Lookup(&RemoteHost{Host: ""})
	if err == nil {
		assert.FailNowf(t, "Expected error", err.Error())
	}
}

func Test_geoProviderIpApi_Lookup_Invalid(t *testing.T) {
	url := "http://localhost:8080/json/%v"

	provider := NewProviderIpApi(url)

	r, err := provider.Lookup(&RemoteHost{Host: "0.0.0.0", IsIp: true})
	if err != nil {
		assert.FailNowf(t, "Error", err.Error())
	}

	if nil == r {
		assert.FailNow(t, "Response is nil")
	}
}

func Test_geoProviderIpApi_Lookup_Basic(t *testing.T) {
	url := "http://localhost:8080/json/%v"

	provider := NewProviderIpApi(url)

	r, err := provider.Lookup(&RemoteHost{Host: "1.1.1.1", IsIp: true})
	if err != nil {
		assert.FailNowf(t, "Error", err.Error())
	}

	if nil == r {
		assert.FailNow(t, "Response is nil")
	}

	expectedIp := "1.1.1.1"
	if r.Ip != expectedIp {
		assert.Failf(t, "Expected IP to be %v but was %v", r.Ip)
	}

	expectedCountry := "Canada"
	if r.Country != expectedCountry {
		assert.Failf(t, "Expected country to be %v but was %v", expectedCountry, r.Country)
	}

	expectedRegion := "QC"
	if r.Region != expectedRegion {
		assert.Failf(t, "Expected region to be %v but was %v", expectedRegion, r.Region)
	}

	expectedCity := "Saint-Leonard"
	if r.City != expectedCity {
		assert.Failf(t, "Expected city to be %v but was %v", expectedCity, r.City)
	}

	expectedOrg := "Videotron Ltee"
	if r.Org != expectedOrg {
		assert.Failf(t, "Expected org to be %v but was %v", expectedOrg, r.Org)
	}

	expectedIsp := "Le Groupe Videotron Ltee"
	if r.Isp != expectedIsp {
		assert.Failf(t, "Expected ISP to be %v but was %v", expectedIsp, r.Isp)
	}
}
