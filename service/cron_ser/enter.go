package cron_ser

import (
	"github.com/robfig/cron/v3"
	"time"
)

// CronInit 定时任务
func CronInit() {

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	Cron.AddFunc("*/10 * * * * *", SyncArticleData)
	Cron.AddFunc("*/10 * * * * *", SyncCommentData)
	Cron.Start()

}
