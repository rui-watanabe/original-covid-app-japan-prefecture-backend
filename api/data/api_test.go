package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"original-covid-app-japan-prefecture-backend/client"
	"testing"
)

func readClientApi() (clientApiData client.ClientApiData) {
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

func TestGetExportApiDate(t *testing.T) {
	clientApi := readClientApi()
	arg := clientApi[0].SubmitDate
	value := getExportApiDate(arg)
	ans := "2021年4月2日"
	if value != ans {
		t.Errorf("%v != %v", value, ans)
	}
}

func TestGetExportApiData(t *testing.T) {
	clientApi := readClientApi()
	exportApi := getExportApiData(clientApi)
	prefName := exportApi.Hokkaido.PrefectureInfo.Name
	ansPrefName := "北海道"
	if prefName != ansPrefName {
		t.Errorf("%v != %v", prefName, ansPrefName)
	}
	count := 1
	ansCount := exportApi.Hokkaido.Hospitalize.Normal
	if count != ansCount {
		t.Errorf("%v != %v", count, ansCount)
	}
}
