package data

import (
	"original-covid-app-japan-prefecture-backend/client"
	"testing"
)

func TestGetExportApiDate(t *testing.T) {
	clientApi := client.ReadClientApi()
	arg := clientApi[0].SubmitDate
	date1 := getExportApiDate(arg)
	date2 := "2021年4月2日"
	if date1 != date2 {
		t.Errorf("%v != %v", date1, date2)
	}
}

func TestGetExportApiData(t *testing.T) {
	clientApi := client.ReadClientApi()
	exportApi := getExportApiData(clientApi)
	prefName1 := exportApi.Hokkaido.PrefectureInfo.Name
	prefName2 := "北海道"
	if prefName1 != prefName2 {
		t.Errorf("%v != %v", prefName1, prefName2)
	}
	count1 := 1
	count2 := exportApi.Hokkaido.Hospitalize.Normal
	if count1 != count2 {
		t.Errorf("%v != %v", count1, count2)
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
