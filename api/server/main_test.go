package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"original-covid-app-japan-prefecture-backend/api/data"
	"testing"
)

func TestStartApiServer(t *testing.T) {
	d, _ := data.InitExportApiData()
	ts := httptest.NewServer(parseURL(topHandler, d))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println(err)
	}

	date1 := d.Date
	date2 := "2021年3月31日"
	if date1 != date2 {
		t.Errorf("%v != %v", date1, date2)
	}

	name1 := d.Data.Hokkaido.PrefectureInfo.Name
	name2 := "北海道"
	if name1 != name2 {
		t.Errorf("%v != %v", name1, name2)
	}

	n1 := d.Data.Hokkaido.OutPatient.Normal
	n2 := 1
	if n1 != n2 {
		t.Errorf("%v != %v", n1, n2)
	}

	l1 := d.Data.Hokkaido.OutPatient.Limit
	l2 := 2
	if l1 != l2 {
		t.Errorf("%v != %v", l1, l2)
	}

	s1 := d.Data.Hokkaido.OutPatient.Stopped
	s2 := 3
	if s1 != s2 {
		t.Errorf("%v != %v", s1, s2)
	}

	ui1 := d.Data.Hokkaido.OutPatient.Unintroduced
	ui2 := 4
	if ui1 != ui2 {
		t.Errorf("%v != %v", ui1, ui2)
	}

	ua1 := d.Data.Hokkaido.OutPatient.Unanswered
	ua2 := 5
	if ua1 != ua2 {
		t.Errorf("%v != %v", ua1, ua2)
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
