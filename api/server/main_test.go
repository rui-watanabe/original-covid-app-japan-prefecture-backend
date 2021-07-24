package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"original-covid-app-japan-prefecture-backend/api/data"
	"os"
	"testing"
)

func TestStartApiServer(t *testing.T) {
	os.Chdir("../..")
	s, err := os.Getwd()
	if err != nil {
		t.Errorf("Get Current Path Error")
		return
	}

	d, _ := data.InitExportApiData(s)
	ts := httptest.NewServer(parseURL(topHandler, d))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}

	err = json.Unmarshal(body, &d)
	if err != nil {
		t.Error(err)
		return
	}

	date1 := d.Date
	date2 := "2021年3月31日"
	if date1 != date2 {
		t.Errorf("%v != %v", date1, date2)
		return
	}

	name1 := d.Data.Hokkaido.PrefectureInfo.Name
	name2 := "北海道"
	if name1 != name2 {
		t.Errorf("%v != %v", name1, name2)
		return
	}

	n1 := d.Data.Hokkaido.OutPatient.Normal
	n2 := 1
	if n1 != n2 {
		t.Errorf("%v != %v", n1, n2)
		return
	}

	l1 := d.Data.Hokkaido.OutPatient.Limit
	l2 := 2
	if l1 != l2 {
		t.Errorf("%v != %v", l1, l2)
		return
	}

	s1 := d.Data.Hokkaido.OutPatient.Stopped
	s2 := 3
	if s1 != s2 {
		t.Errorf("%v != %v", s1, s2)
		return
	}

	ui1 := d.Data.Hokkaido.OutPatient.Unintroduced
	ui2 := 4
	if ui1 != ui2 {
		t.Errorf("%v != %v", ui1, ui2)
		return
	}

	ua1 := d.Data.Hokkaido.OutPatient.Unanswered
	ua2 := 5
	if ua1 != ua2 {
		t.Errorf("%v != %v", ua1, ua2)
		return
	}

	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}
