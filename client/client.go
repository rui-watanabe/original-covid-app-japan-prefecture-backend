package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient(c *http.Client) Client {
	return Client{c}
}

type ClientApiData []struct {
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

func (c Client) FetchClientApiData() (clientApiData ClientApiData) {
	req, err := http.NewRequest("GET", "https://opendata.corona.go.jp/api/covid19DailySurvey", nil)
	if err != nil {
		fmt.Println(err)
	}
	res, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &clientApiData)
	if err != nil {
		fmt.Println(err)
	}
	return clientApiData
}

func ReadClientApi() (clientApiData ClientApiData) {
	jsonFile, err := ioutil.ReadFile("../../client/clientApi.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile, &clientApiData)
	if err != nil {
		fmt.Println(err)
	}
	return
}
