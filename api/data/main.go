package data

import (
	"original-covid-app-japan-prefecture-backend/client"
)

func MakeApiJson(clientApi client.ClientApi) ExportApi {
	exportApi := ExportApi{}
	exportApi.Date = getExportApiDate(clientApi[0].SubmitDate)
	exportApi.Data = getExportApiData(clientApi)
	return exportApi
}
