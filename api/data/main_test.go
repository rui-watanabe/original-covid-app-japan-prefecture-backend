package data

import (
	"original-covid-app-japan-prefecture-backend/client"
	"testing"
)

func TestMakeApiJson(t *testing.T) {
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

	data := MakeApiJson(c)

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
