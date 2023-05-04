package new_api

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/utils/requests"
	"io"
	"time"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}
type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewData struct {
	Index    string `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

type NewsResponse struct {
	Code int       `json:"code"`
	Data []NewData `json:"data"`
	Msg  string    `json:"msg"`
}

const newAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

func (NewsApi) NewsListView(c *gin.Context) {
	var cr params
	var headers header
	err := c.ShouldBindJSON(&cr)
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	if cr.Size == 0 {
		cr.Size = 1
	}

	httpResponse, err := requests.Post(newAPI, cr, structs.Map(headers), timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	var response NewsResponse
	byteData, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}
	fmt.Println(response.Data)
	res.OkWithData(response.Data, c)
	return
}
