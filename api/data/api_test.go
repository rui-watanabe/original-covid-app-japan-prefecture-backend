package data

import "testing"

func TestGetExportApiDate(t *testing.T) {
	arg := "2021-07-03"
	value := getExportApiDate(arg)
	ans := "2021年7月3日"
	if value != ans {
		t.Errorf("%v != %v", value, ans)
	}
}
