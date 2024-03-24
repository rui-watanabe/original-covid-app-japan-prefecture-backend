package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient(c *http.Client) Client {
	return Client{c}
}

type ClientApi []struct {
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

func (c Client) FetchClientApi() (clientApi ClientApi) {
	url := "https://opendata.corona.go.jp/api/covid19DailySurvey"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := c.Do(req)
	if err != nil {
		err = json.Unmarshal(clientData, &clientApi)
		if err != nil {
			log.Fatalln(err, "a")
		} else {
			return clientApi
		}
		log.Fatalln(err, clientApi)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err, "c")
	}
	err = json.Unmarshal(body, &clientApi)
	if err != nil {
		log.Fatalln(err, "d")
	}
	return clientApi
}
