package data

func JudgeExportApiDataFieldNumber(prefecture string) (name string, num int) {
	name = prefecture
	switch name {
	case "北海道":
		num = 0
	case "青森県":
		num = 1
	case "岩手県":
		num = 2
	case "宮城県":
		num = 3
	case "秋田県":
		num = 4
	case "山形県":
		num = 5
	case "福島県":
		num = 6
	case "茨城県":
		num = 7
	case "栃木県":
		num = 8
	case "群馬県":
		num = 9
	case "埼玉県":
		num = 10
	case "千葉県":
		num = 11
	case "東京都":
		num = 12
	case "神奈川県":
		num = 13
	case "新潟県":
		num = 14
	case "富山県":
		num = 15
	case "石川県":
		num = 16
	case "福井県":
		num = 17
	case "山梨県":
		num = 18
	case "長野県":
		num = 19
	case "岐阜県":
		num = 20
	case "静岡県":
		num = 21
	case "愛知県":
		num = 22
	case "三重県":
		num = 23
	case "滋賀県":
		num = 24
	case "京都府":
		num = 25
	case "大阪府":
		num = 26
	case "兵庫県":
		num = 27
	case "奈良県":
		num = 28
	case "和歌山県":
		num = 29
	case "鳥取県":
		num = 30
	case "島根県":
		num = 31
	case "岡山県":
		num = 32
	case "広島県":
		num = 33
	case "山口県":
		num = 34
	case "徳島県":
		num = 35
	case "香川県":
		num = 36
	case "愛媛県":
		num = 37
	case "高知県":
		num = 38
	case "福岡県":
		num = 39
	case "佐賀県":
		num = 40
	case "長崎県":
		num = 41
	case "熊本県":
		num = 42
	case "大分県":
		num = 43
	case "宮崎県":
		num = 44
	case "鹿児島県":
		num = 45
	case "沖縄県":
		num = 46
	}
	return
}

func JudgeExportApiInfoFieledNumber(facility string) (num int) {
	switch facility {
	case "入院":
		num = 1
	case "外来":
		num = 2
	case "救急":
		num = 3
	}
	return num
}

func JudgeExportApiCountFieledNumber(ans string) (num int) {
	switch ans {
	case "通常":
		num = 0
	case "制限":
		num = 1
	case "停止":
		num = 2
	case "設置なし":
		num = 3
	case "未回答":
		num = 4
	}
	return num
}
