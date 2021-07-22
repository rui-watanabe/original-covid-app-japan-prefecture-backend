package data

import (
	"original-covid-app-japan-prefecture-backend/client"
	"reflect"
)

func getExportApiDate(clientApi client.ClientApiData) string {
	date := clientApi[0].SubmitDate
	str := date[0:4] + "年" + date[5:7] + "月" + date[8:10] + "日"
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
