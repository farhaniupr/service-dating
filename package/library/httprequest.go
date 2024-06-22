package library

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func HttpRequestAPI(c context.Context, method string, url string,
	req map[string]interface{}, header map[string]string, target interface{}) (int, string, error) {

	jsonValue, err := json.Marshal(req)
	if err != nil {
		return 500, "", err
	}

	bytejson := bytes.NewBuffer(jsonValue)

	client := &http.Client{Timeout: 300 * time.Second}
	req_url, err := http.NewRequest(method, url, bytejson)

	if err != nil {
		return 500, "", err
	}
	for key, value := range header {
		req_url.Header.Add(key, value)
	}
	req_url.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req_url)
	if err != nil {
		return 500, "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 500, "", err
	}

	defer res.Body.Close()

	return res.StatusCode, string(body), json.Unmarshal([]byte(body), &target)
}
