package monitlog

import (
	"encoding/json"
	"bytes"
	"net/http"
	"net/url"
)

type LogMesg struct {
	Id    string`json:"id"`
	Level int `json:"level"` // level = 0 ERROR, 1:INFO
	Mesg  string `json:"mesg"`
}

var ERROR = 0
var INFO = 1

// post to Monitor
// curl -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d "{\"id\":\"af387ed3-f4bc-11e6-8d73-507b9dcc7a61\",\"level\":1,\"mesg\":\"Error?\" }" http:/<host>/log
func PostToMonitor(log LogMesg, API_URL string) {
	jsonString, _ := json.Marshal(log)
	client := &http.Client{
	}
	//data := url.Values{}
	//data.Set("params", PARAMS)
	req, _ := http.NewRequest("POST", API_URL, bytes.NewBufferString(string(jsonString)))
	client.Do(req)

}

// post status Monitor
func PostStatusToMonitor(serviceid string, status string, API_URL string) {
	client := &http.Client{
	}
	data := url.Values{}
	data.Set("service_id", serviceid)
	data.Set("log", status)
	req, _ := http.NewRequest("POST", API_URL, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client.Do(req)

}
