package server

import (
	"encoding/json"
	"github.com/rs/cors"
	"log"
	"net/http"
	"original-covid-app-japan-prefecture-backend/api/data"
	// "original-covid-app-japan-prefecture-backend/config"
)

func StartApiServer(exportApi data.ExportApi) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", parseURL(topHandler, exportApi))

	// CORS レスポンスヘッダーの追加
	c := cors.Default()
	handler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":5000", handler))
}

func parseURL(fn func(http.ResponseWriter, *http.Request, data.ExportApi), data data.ExportApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, data)
	}
}

func topHandler(w http.ResponseWriter, r *http.Request, exportApi data.ExportApi) {
	json.NewEncoder(w).Encode(exportApi)
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
