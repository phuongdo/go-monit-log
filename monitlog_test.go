package monitlog

import "testing"


//curl -H "Accept: application/json" -H "Content-Type: application/json"  -X POST --data '@req.json' API_URI
func TestLogging(t *testing.T) {
	API_URL := "http://localhost:9999/log"
	serviceId := "af387d3c-f4bc-11e6-8d6e-507b9dcc7a61"
	logErro := LogMesg{serviceId, ERROR, "your error content"}
	PostToMonitor(logErro, API_URL)
	logInfo := LogMesg{serviceId, ERROR, "your info content"}
	PostToMonitor(logInfo, API_URL)
}

func TestStatusLogging(t *testing.T) {
	API_URL := "http://localhost:9999"
	serviceId := "af387ed3-f4bc-11e6-8d73-507b9dcc7a61"
	status := "running"
	PostStatusToMonitor(serviceId, status, API_URL)
}
