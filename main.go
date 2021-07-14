package main

import (
	"fmt"
	"net/http"
	"original-covid-app-japan-prefecture-backend/client"
)

func main() {
	apiClient := client.NewClient(http.DefaultClient)
	extraApiData := apiClient.FetchApiData()
	fmt.Println(extraApiData)
}
