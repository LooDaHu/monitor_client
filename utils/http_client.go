package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ServerAddress        = "servername.com"
	KUploadSystemInfoURL = "/system/add"
	AppCode              = "1234567890"
)

var client *http.Client

func InitHTTPClient() {
	client = &http.Client{}
}

func HttpClient(data interface{}) error {

	jsonData, jErr := json.Marshal(data)
	if jErr != nil {
		SugarLogger.Error("Json Marshal Error @HttpClient", jErr)
		return MarshalError
	}
	req, err := http.NewRequest("POST",
		strings.Join([]string{"https://", ServerAddress, KUploadSystemInfoURL}, ""),
		bytes.NewReader(jsonData))
	if err != nil {
		SugarLogger.Error("HTTP Request Error @HttpClient", err)
		return HTTPError
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", strings.Join([]string{"APPCODE", AppCode}, " "))

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		SugarLogger.Error("HTTP Request Error @HttpClient", err)
		return HTTPError
	}

	fmt.Println(string(body))
	return nil
}
