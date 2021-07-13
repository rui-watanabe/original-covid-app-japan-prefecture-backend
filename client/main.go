package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient(c *http.Client) Client {
	return Client{c}
}

type ExtraApiData []struct {
	FacilityID   string `json:"facilityId"`
	FacilityName string `json:"facilityName"`
	ZipCode      string `json:"zipCode"`
	PrefName     string `json:"prefName"`
	FacilityAddr string `json:"facilityAddr"`
	FacilityTel  string `json:"facilityTel"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	SubmitDate   string `json:"submitDate"`
	FacilityType string `json:"facilityType"`
	AnsType      string `json:"ansType"`
	LocalGovCode string `json:"localGovCode"`
	CityName     string `json:"cityName"`
	FacilityCode string `json:"facilityCode"`
}

func (c Client) FetchApiData() (extraApiData *ExtraApiData) {
	req, _ := http.NewRequest("GET", "https://opendata.corona.go.jp/api/covid19DailySurvey", nil)
	res, _ := c.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &extraApiData)
	return extraApiData
}
