package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"original-covid-app-japan-prefecture-backend/api/data"
	"original-covid-app-japan-prefecture-backend/client"
	"testing"
)

func TestStartApiServer(t *testing.T) {
	c := client.ClientApi{{
		FacilityID:   "11111111",
		FacilityName: "病院",
		ZipCode:      "〒000-0000",
		PrefName:     "北海道",
		FacilityAddr: "金沢市桜町00の00",
		FacilityTel:  "1111111111",
		Latitude:     "11.111111",
		Longitude:    "111.111111",
		SubmitDate:   "2021-04-02",
		FacilityType: "入院",
		AnsType:      "通常",
		LocalGovCode: "111111",
		CityName:     "金沢市",
		FacilityCode: "1111111111",
	}}

	d := data.MakeApiJson(c)
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
	date2 := "2021年4月2日"
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

	n1 := d.Data.Hokkaido.Hospitalize.Normal
	n2 := 1
	if n1 != n2 {
		t.Errorf("%v != %v", n1, n2)
		return
	}

	l1 := d.Data.Hokkaido.Hospitalize.Limit
	l2 := 0
	if l1 != l2 {
		t.Errorf("%v != %v", l1, l2)
		return
	}

	s1 := d.Data.Hokkaido.Hospitalize.Stopped
	s2 := 0
	if s1 != s2 {
		t.Errorf("%v != %v", s1, s2)
		return
	}

	ui1 := d.Data.Hokkaido.Hospitalize.Unintroduced
	ui2 := 0
	if ui1 != ui2 {
		t.Errorf("%v != %v", ui1, ui2)
		return
	}

	ua1 := d.Data.Hokkaido.Hospitalize.Unanswered
	ua2 := 0
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
