package main

import (
	"log"
	"original-covid-app-japan-prefecture-backend/api"
	"original-covid-app-japan-prefecture-backend/client"
)

func main() {
	//クライアント
	clientApi := client.GetApiData()
	//APIサーバー
	api.StartServer(clientApi)
}
