package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

type Params struct {
	ID string `json:"id"`
}

type NewsResponse struct {
	Code int `json:"code"`
	Data []struct {
		Index    string `json:"index"`
		Title    string `json:"title"`
		HotValue string `json:"hotValue"`
		Link     string `json:"link"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func main() {
	var params = Params{
		ID: "mproPpoq6O",
	}
	reqParam, _ := json.Marshal(params)
	reqBody := strings.NewReader(string(reqParam))
	url := "https://api.codelife.cc/api/top/list"
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		logrus.Error(err)
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("version", "1.3.35")
	httpReq.Header.Add("signaturekey", "U2FsdGVkX19R+FnzVgqWm4fp/6b4yAXdeBcmT3Uwods=")

	client := http.Client{
		Timeout: 2 * time.Second,
	}
	// DO: HTTP请求
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		logrus.Error(err)
		return
	}
	ByteData, err := io.ReadAll(httpRsp.Body)

	var response NewsResponse
	err = json.Unmarshal(ByteData, &response)
	if err != nil {
		logrus.Error(err)
		return
	}
	if response.Code != 200 {
		logrus.Error(response.Msg)
		return
	}
	fmt.Println(response)
}
