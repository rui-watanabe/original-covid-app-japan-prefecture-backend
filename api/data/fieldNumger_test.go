package data

import "testing"

func TestJudgeExportApiDataFieldNumber(t *testing.T) {
	n := 0
	v := JudgeExportApiDataFieldNumber("北海道")
	if n != v {
		t.Errorf("%v != %v", n, v)
	}
}

func TestJudgeExportApiInfoFieledNumber(t *testing.T) {
	n := 1
	v := JudgeExportApiInfoFieledNumber("入院")
	if n != v {
		t.Errorf("%v != %v", n, v)
	}
}

func TestJudgeExportApiCountFieledNumber(t *testing.T) {
	n := 0
	v := JudgeExportApiCountFieledNumber("通常")
	if n != v {
		t.Errorf("%v != %v", n, v)
	}
}
