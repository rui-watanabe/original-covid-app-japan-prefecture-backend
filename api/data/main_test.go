package data

import (
	"original-covid-app-japan-prefecture-backend/client"
	"testing"
)

func TestMakeApiJson(t *testing.T) {
	arg := client.ReadClientApi()
	data := MakeApiJson(arg)

	date1 := data.Date
	date2 := "2021年4月2日"
	if date1 != date2 {
		t.Errorf("%v != %v", date1, date2)
	}

	prefName1 := data.Data.Hokkaido.PrefectureInfo.Name
	prefName2 := "北海道"
	if prefName1 != prefName2 {
		t.Errorf("%v != %v", prefName1, prefName2)
	}
	count1 := 1
	count2 := data.Data.Hokkaido.Hospitalize.Normal
	if count1 != count2 {
		t.Errorf("%v != %v", count1, count2)
	}
}
