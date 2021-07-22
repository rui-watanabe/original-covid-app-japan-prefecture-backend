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

func TestInitExportApiData(t *testing.T) {
	data := initExportApiData()
	n1 := data.Hokkaido.OutPatient.Normal
	n2 := 1
	if n1 != n2 {
		t.Errorf("%v != %v", n1, n2)
	}

	l1 := data.Hokkaido.OutPatient.Limit
	l2 := 2
	if l1 != l2 {
		t.Errorf("%v != %v", l1, l2)
	}

	s1 := data.Hokkaido.OutPatient.Stopped
	s2 := 3
	if s1 != s2 {
		t.Errorf("%v != %v", s1, s2)
	}

	ui1 := data.Hokkaido.OutPatient.Unintroduced
	ui2 := 4
	if ui1 != ui2 {
		t.Errorf("%v != %v", ui1, ui2)
	}

	ua1 := data.Hokkaido.OutPatient.Unanswered
	ua2 := 5
	if ua1 != ua2 {
		t.Errorf("%v != %v", ua1, ua2)
	}
}
