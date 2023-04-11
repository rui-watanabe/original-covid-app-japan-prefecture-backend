package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"original-covid-app-japan-prefecture-backend/api/data"
	// "original-covid-app-japan-prefecture-backend/config"
)

func StartApiServer(exportApi data.ExportApi) {
	http.HandleFunc("/", parseURL(topHandler, exportApi))
	print("hogehoge")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func parseURL(fn func(http.ResponseWriter, *http.Request, data.ExportApi), data data.ExportApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, data)
	}
}

func topHandler(w http.ResponseWriter, r *http.Request, exportApi data.ExportApi) {
	json.NewEncoder(w).Encode(exportApi)
	file, _ := json.MarshalIndent(exportApi, "", " ")
	_ = ioutil.WriteFile("opt/build/repo/data.json", file, 0644)
}
