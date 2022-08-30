package corn

import (
	"github.com/aecra/covid/report"
	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()
	// 考虑是否将两者分开
	c.AddFunc("57 10 8,12 * * *", report.ReportAllClock)
	c.AddFunc("39 27 7 * * *", report.ReportAllHealth)
	c.Start()
}
