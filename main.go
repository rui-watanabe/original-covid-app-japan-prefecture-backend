package main

import (
	"fmt"
	"original-covid-app-japan-prefecture-backend/api"
	"original-covid-app-japan-prefecture-backend/client"
	"original-covid-app-japan-prefecture-backend/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.LogFile)
	//クライアント
	clientApi := client.GetApiData()
	//APIサーバー
	api.StartServer(clientApi)
}
