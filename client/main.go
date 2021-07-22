package client

import (
	"net/http"
)

func GetApiData() (clientApiData ClientApiData) {
	apiClient := NewClient(http.DefaultClient)
	clientApiData = apiClient.FetchClientApiData()
	return
}
