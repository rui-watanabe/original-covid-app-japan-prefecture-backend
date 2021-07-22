package server

import (
	"encoding/json"
	"net/http"
	"original-covid-app-japan-prefecture-backend/api/data"
)

func StartApiServer(exportApiData data.ExportApi) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(exportApiData)
	})
	http.ListenAndServe(":8000", nil)
}
