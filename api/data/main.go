package data

import (
	"original-covid-app-japan-prefecture-backend/client"
)

func MakeApiJson(clientApi client.ClientApiData) ExportApi {
	exportApi := ExportApi{}
	exportApi.Date = getExportApiDate(clientApi)
	exportApi.Data = getExportApiData(clientApi)
	return exportApi
}
