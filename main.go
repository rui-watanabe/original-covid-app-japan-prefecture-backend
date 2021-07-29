package main

import (
	"original-covid-app-japan-prefecture-backend/api"
	"original-covid-app-japan-prefecture-backend/client"
	"original-covid-app-japan-prefecture-backend/config"
	"original-covid-app-japan-prefecture-backend/utils"
)

func init() {
	c := config.LoadConfig()
	utils.LoggingSettings(c.LogFile)
}

func main() {
	//クライアント
	clientApi := client.GetApiData()
	//APIサーバー
	api.StartServer(clientApi)
}
