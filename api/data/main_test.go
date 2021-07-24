package data

import (
	"encoding/json"
	"io/ioutil"
	"original-covid-app-japan-prefecture-backend/client"
	"os"
	"testing"
)

func TestMakeApiJson(t *testing.T) {
	s, err := os.Getwd()
	if err != nil {
		t.Errorf("Get Current Path Error")
		return
	}
	clientApi := client.ClientApi{}

	jsonFile, err := ioutil.ReadFile(s + "/client/clientApi.json")
	if err != nil {
		t.Error(err)
		return
	}
	err = json.Unmarshal(jsonFile, &clientApi)
	if err != nil {
		t.Error(err)
		return
	}

	data := MakeApiJson(clientApi)

	date1 := data.Date
	date2 := "2021年4月2日"
	if date1 != date2 {
		t.Errorf("%v != %v", date1, date2)
		return
	}

	prefName1 := data.Data.Hokkaido.PrefectureInfo.Name
	prefName2 := "北海道"
	if prefName1 != prefName2 {
		t.Errorf("%v != %v", prefName1, prefName2)
		return
	}
	count1 := 1
	count2 := data.Data.Hokkaido.Hospitalize.Normal
	if count1 != count2 {
		t.Errorf("%v != %v", count1, count2)
		return
	}
}
