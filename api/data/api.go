package data

import (
	"fmt"
	"original-covid-app-japan-prefecture-backend/client"
	"reflect"
	"strconv"
)

func getExportApiDate(date string) string {
	y, err := strconv.Atoi(date[0:4])
	if err != nil {
		fmt.Println(err)
	}
	m, err := strconv.Atoi(date[5:7])
	if err != nil {
		fmt.Println(err)
	}
	d, err := strconv.Atoi(date[8:10])
	if err != nil {
		fmt.Println(err)
	}
	str := strconv.Itoa(y) + "年" + strconv.Itoa(m) + "月" + strconv.Itoa(d) + "日"
	return str
}

func getExportApiData(clientApi client.ClientApiData) ExportApiData {
	data := ExportApiData{}
	uv := reflect.ValueOf(&data).Elem()
	for _, value := range clientApi {
		prefName, prefNum := JudgeExportApiDataFieldNumber(value.PrefName)
		infoNum := JudgeExportApiInfoFieledNumber(value.FacilityType)
		countNum := JudgeExportApiCountFieledNumber(value.AnsType)
		count := uv.Field(prefNum).Field(infoNum).Field(countNum)
		sum := count.Int() + 1
		uv.Field(prefNum).Field(0).Field(0).SetString(prefName)
		uv.Field(prefNum).Field(infoNum).Field(countNum).SetInt(sum)
	}
	return data
}
