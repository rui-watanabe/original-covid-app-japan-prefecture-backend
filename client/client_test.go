package client

import (
	"net/http"
	"testing"
)

func TestFetchClientApi(t *testing.T) {
	apiClient := NewClient(http.DefaultClient)
	clientApi := apiClient.FetchClientApi()
	if len(clientApi) == 0 {
		t.Error("Not get client api")
	}
}

func TestReadClientApi(t *testing.T) {
	clientApi := ReadClientApi()
	if len(clientApi) == 0 {
		t.Error("Not get init client api")
	}
}
