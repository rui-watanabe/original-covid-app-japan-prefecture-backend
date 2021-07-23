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

func InitExportApiData() (exportApi ExportApi, data ExportApiData) {
	s := os.Getenv("GOPATH")
	if s == "" {
		log.Fatalln("Not GOPATH")
	}
	jsonFile, err := ioutil.ReadFile(s + "/src/original-covid-app-japan-prefecture-backend/api/data/exportApi.json")
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
	_, data = InitExportApiData()
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
