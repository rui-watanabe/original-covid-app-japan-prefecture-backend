package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"original-covid-app-japan-prefecture-backend/client"
	"os"
	"reflect"
	"strconv"
)

func InitExportApiData(path string) (exportApi ExportApi, data ExportApiData) {
	jsonFile, err := ioutil.ReadFile(path + "/api/data/exportApi.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(jsonFile, &exportApi)
	if err != nil {
		log.Fatalln(err)
	}
	data = exportApi.Data
	return
}

func getExportApiDate(date string) string {
	y, err := strconv.Atoi(date[0:4])
	if err != nil {
		log.Fatalln(err)
	}
	m, err := strconv.Atoi(date[5:7])
	if err != nil {
		log.Fatalln(err)
	}
	d, err := strconv.Atoi(date[8:10])
	if err != nil {
		log.Fatalln(err)
	}
	str := strconv.Itoa(y) + "年" + strconv.Itoa(m) + "月" + strconv.Itoa(d) + "日"
	return str
}

func getExportApiData(clientApi client.ClientApi) (data ExportApiData) {
	s, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get Current Path Error")
	}
	_, data = InitExportApiData(s)
	uv := reflect.ValueOf(&data).Elem()
	for _, value := range clientApi {
		prefNum := JudgeExportApiDataFieldNumber(value.PrefName)
		infoNum := JudgeExportApiInfoFieledNumber(value.FacilityType)
		countNum := JudgeExportApiCountFieledNumber(value.AnsType)
		count := uv.Field(prefNum).Field(infoNum).Field(countNum)
		sum := count.Int() + 1
		uv.Field(prefNum).Field(infoNum).Field(countNum).SetInt(sum)
	}
	return data
}
