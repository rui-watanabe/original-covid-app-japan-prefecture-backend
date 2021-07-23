package client

import (
	"net/http"
)

func GetApiData() (clientApi ClientApi) {
	apiClient := NewClient(http.DefaultClient)
	clientApi = apiClient.FetchClientApi()
	return
}
