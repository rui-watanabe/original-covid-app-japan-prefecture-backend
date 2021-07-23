package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	req, err := http.NewRequest("GET", "https://opendata.corona.go.jp/api/covid19DailySurvey", nil)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &clientApi)
	if err != nil {
		log.Fatalln(err)
	}
	return clientApi
}

func ReadClientApi() (clientApi ClientApi) {
	s := os.Getenv("GOPATH")
	if s == "" {
		log.Fatalln("Not GOPATH")
	}
	jsonFile, err := ioutil.ReadFile(s + "/src/original-covid-app-japan-prefecture-backend/client/clientApi.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(jsonFile, &clientApi)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
