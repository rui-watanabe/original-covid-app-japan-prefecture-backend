package main

import (
	"original-covid-app-japan-prefecture-backend/api"
	"original-covid-app-japan-prefecture-backend/client"
)

func main() {
	//クライアント
	clientApiData := client.GetApiData()
	//APIサーバー
	api.StartServer(clientApiData)
}
