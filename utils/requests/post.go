package requests

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Post(url string, data any, headers map[string]interface{}, timeout time.Duration) (resp *http.Response, err error) {
	reqParam, _ := json.Marshal(data)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		logrus.Error(err)
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		switch v := val.(type) {
		case string:
			httpReq.Header.Add(key, v)
		case int:
			httpReq.Header.Add(key, strconv.Itoa(v))
		}
	}
	client := http.Client{
		Timeout: timeout,
	}
	// DO: HTTP请求
	httpResp, err := client.Do(httpReq)
	return httpResp, err
}
