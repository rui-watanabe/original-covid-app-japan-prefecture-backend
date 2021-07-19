package main

import (
	"fmt"
	"net/http"
	"original-covid-app-japan-prefecture-backend/api"
	"original-covid-app-japan-prefecture-backend/client"
)

func main() {
	//クライアント
	apiClient := client.NewClient(http.DefaultClient)
	clientApiData := apiClient.FetchClientApiData()

	//APIデータ作成
	exportApiData := api.MakeApiJson(clientApiData)
	fmt.Println(exportApiData)
}
