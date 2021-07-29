package data

import (
	"log"
	"original-covid-app-japan-prefecture-backend/client"
	"reflect"
	"strconv"
)

var prefNameArray = [47]string{"北海道", "青森県", "岩手県", "宮城県", "秋田県", "山形県", "福島県",
	"茨城県", "栃木県", "群馬県", "埼玉県", "千葉県", "東京都", "神奈川県",
	"新潟県", "富山県", "石川県", "福井県", "山梨県", "長野県", "岐阜県",
	"静岡県", "愛知県", "三重県", "滋賀県", "京都府", "大阪府", "兵庫県",
	"奈良県", "和歌山県", "鳥取県", "島根県", "岡山県", "広島県", "山口県",
	"徳島県", "香川県", "愛媛県", "高知県", "福岡県", "佐賀県", "長崎県",
	"熊本県", "大分県", "宮崎県", "鹿児島県", "沖縄県"}

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
	uv := reflect.ValueOf(&data).Elem()
	for index, pref := range prefNameArray {
		uv.Field(index).Field(0).Field(0).SetString(pref)
	}
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
