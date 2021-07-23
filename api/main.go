package api

import (
	"original-covid-app-japan-prefecture-backend/api/data"
	"original-covid-app-japan-prefecture-backend/api/server"
	"original-covid-app-japan-prefecture-backend/client"
)

func StartServer(clientApi client.ClientApi) {
	exportApiData := data.MakeApiJson(clientApi)
	server.StartApiServer(exportApiData)
}
