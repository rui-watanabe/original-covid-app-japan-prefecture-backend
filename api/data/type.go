package data

type ExportApiPrefectureInfo struct {
	Name string `json:"name"`
}

type ExportApiCount struct {
	Normal       int `json:"normal"`
	Limit        int `json:"limit"`
	Stopped      int `json:"stopped"`
	Unintroduced int `json:"unintroduced"`
	Unanswered   int `json:"unanswered"`
}

type ExportApiInfo struct {
	PrefectureInfo ExportApiPrefectureInfo `json:"prefectureInfo"`
	Hospitalize    ExportApiCount          `json:"hospitalize"`
	OutPatient     ExportApiCount          `json:"outPatient"`
	Emergency      ExportApiCount          `json:"emergency"`
}

type ExportApiData struct {
	Hokkaido  ExportApiInfo `json:"hokkaido"`
	Aomori    ExportApiInfo `json:"aomori"`
	Iwate     ExportApiInfo `json:"iwate"`
	Miyagi    ExportApiInfo `json:"miyagi"`
	Akita     ExportApiInfo `json:"akita"`
	Yamagata  ExportApiInfo `json:"yamagata"`
	Fukushima ExportApiInfo `json:"fukushima"`
	Ibaraki   ExportApiInfo `json:"ibaraki"`
	Tochigi   ExportApiInfo `json:"tochigi"`
	Gunma     ExportApiInfo `json:"gunma"`
	Saitama   ExportApiInfo `json:"saitama"`
	Chiba     ExportApiInfo `json:"chiba"`
	Tokyo     ExportApiInfo `json:"tokyo"`
	Kanagawa  ExportApiInfo `json:"kanagawa"`
	Niigata   ExportApiInfo `json:"niigata"`
	Toyama    ExportApiInfo `json:"toyama"`
	Ishikawa  ExportApiInfo `json:"ishikawa"`
	Fukui     ExportApiInfo `json:"fukui"`
	Yamanashi ExportApiInfo `json:"yamanashi"`
	Nagano    ExportApiInfo `json:"nagano"`
	Gifu      ExportApiInfo `json:"gifu"`
	Shizuoka  ExportApiInfo `json:"shizuoka"`
	Aichi     ExportApiInfo `json:"aichi"`
	Mie       ExportApiInfo `json:"mie"`
	Shiga     ExportApiInfo `json:"shiga"`
	Kyoto     ExportApiInfo `json:"kyoto"`
	Osaka     ExportApiInfo `json:"osaka"`
	Hyogo     ExportApiInfo `json:"hyogo"`
	Nara      ExportApiInfo `json:"nara"`
	Wakayama  ExportApiInfo `json:"wakayama"`
	Tottori   ExportApiInfo `json:"tottori"`
	Shimane   ExportApiInfo `json:"shimane"`
	Okayama   ExportApiInfo `json:"okayama"`
	Hiroshima ExportApiInfo `json:"hiroshima"`
	Yamaguchi ExportApiInfo `json:"yamaguchi"`
	Tokushima ExportApiInfo `json:"tokushima"`
	Kagawa    ExportApiInfo `json:"kagawa"`
	Ehime     ExportApiInfo `json:"ehime"`
	Kochi     ExportApiInfo `json:"kochi"`
	Fukuoka   ExportApiInfo `json:"fukuoka"`
	Saga      ExportApiInfo `json:"saga"`
	Nagasaki  ExportApiInfo `json:"nagasaki"`
	Kumamoto  ExportApiInfo `json:"kumamoto"`
	Oita      ExportApiInfo `json:"oita"`
	Miyazaki  ExportApiInfo `json:"miyazaki"`
	Kagoshima ExportApiInfo `json:"kagoshima"`
	Okinawa   ExportApiInfo `json:"okinawa"`
}

type ExportApi struct {
	Date              string        `json:"date"`
	CurrentPrefecture string        `json:"currentPrefecture"`
	Data              ExportApiData `json:"data"`
}
