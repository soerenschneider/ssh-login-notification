package geo

type IpGeoInfo struct {
	Ip 			string `json:"ip,omitempty"`
	City		string `json:"city,omitempty"`
	Region	 	string `json:"region,omitempty"`
	Country		string `json:"country,omitempty"`
	Isp 		string `json:isp,omitempty`
	Org 		string `json:org,omitempty`
}