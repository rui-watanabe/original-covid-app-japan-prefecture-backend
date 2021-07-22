package data

import "testing"

func TestJudgeExportApiDataFieldNumber(t *testing.T) {
	n := 0
	v := JudgeExportApiDataFieldNumber("北海道")

	if n != v {
		t.Errorf("%v != %v", n, v)
	}
}
