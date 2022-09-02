package corn

import (
	"github.com/aecra/covid/report"
	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()
	if _, err := c.AddFunc("13 8,12 * * *", report.ReportAllClock); err != nil {
		panic(err)
	}
	if _, err := c.AddFunc("39 7 * * *", report.ReportAllHealth); err != nil {
		panic(err)
	}
	c.Start()
	select {}
}
