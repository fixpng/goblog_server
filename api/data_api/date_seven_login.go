package data_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
}

func (DateApi) SevenLogin(c *gin.Context) {
	var loginDateCount, signDateCount []DateCount

	/*	select date_format(created_at, '%Y-%m-%d') as date, count(id) as count
		from login_data_models
		where date_sub(curdate(), interval 7 day) <= created_at
		group by date;	*/
	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(), interval 7 day)<= created_at").
		Select("date_format(created_at,'%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDateCount)

	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(), interval 7 day)<= created_at").
		Select("date_format(created_at,'%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&signDateCount)

	var loginDateCountMap = map[string]int{}
	var signDateCountMap = map[string]int{}
	var loginCountList, signCountList []int
	now := time.Now()
	for _, d := range loginDateCount {
		loginDateCountMap[d.Date] = d.Count
	}
	for _, d := range signDateCount {
		signDateCountMap[d.Date] = d.Count
	}
	var dateList []string
	// 7 天内
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginDateCountMap[day])
		signCountList = append(signCountList, signDateCountMap[day])
	}

	res.OkWithData(DateCountResponse{
		DateList:  dateList,
		LoginData: loginCountList,
		SignData:  signCountList,
	}, c)
}
