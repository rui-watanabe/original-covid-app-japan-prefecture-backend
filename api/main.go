package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"original-covid-app-japan-prefecture-backend/client"
	"reflect"
)

func MakeApiJson(clientApi client.ClientApiData) string {
	j := setExportApi(clientApi)
	return string(j)
}

func setExportApi(clientApi client.ClientApiData) []byte {
	exportApi := initExportApi()
	exportApi.Date = getExportApiDate(clientApi)
	exportApi.Data = getExportApiData(clientApi)
	apiJson, err := json.Marshal(exportApi)
	if err != nil {
		fmt.Println(err)
	}
	return apiJson
}

func initExportApi() ExportApi {
	initApi := ExportApi{}
	jsonFile, err := ioutil.ReadFile("./api/api.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile, &initApi)
	if err != nil {
		fmt.Println(err)
	}
	return initApi
}

func getExportApiDate(clientApi client.ClientApiData) string {
	date := clientApi[0].SubmitDate
	str := date[0:4] + "年" + date[5:7] + "月" + date[8:10] + "日"
	return str
}

func getExportApiData(clientApi client.ClientApiData) ExportApiData {
	data := ExportApiData{}
	uv := reflect.ValueOf(&data).Elem()
	for _, value := range clientApi {
		dataNum := JudgeExportApiDataFieldNumber(value.PrefName)
		infoNum := JudgeExportApiInfoFieledNumber(value.FacilityType)
		countNum := JudgeExportApiCountFieledNumber(value.AnsType)
		count := uv.Field(dataNum).Field(infoNum).Field(countNum)
		sum := count.Int() + 1
		uv.Field(dataNum).Field(infoNum).Field(countNum).SetInt(sum)
	}
	return data
}
