package api

import (
	"original-covid-app-japan-prefecture-backend/api/data"
	"original-covid-app-japan-prefecture-backend/api/server"
	"original-covid-app-japan-prefecture-backend/client"
)

func StartServer(clientApiData client.ClientApiData) {
	exportApiData := data.MakeApiJson(clientApiData)
	server.StartApiServer(exportApiData)
}
