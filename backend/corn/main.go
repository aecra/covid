package corn

import (
	"log"
	"time"

	"github.com/aecra/covid/report"
	"github.com/robfig/cron/v3"
)

func Start() {
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(nyc))
	if _, err := c.AddFunc("13 8,12 * * *", report.ReportAllClock); err != nil {
		log.Fatal(err)
	}
	if _, err := c.AddFunc("39 7 * * *", report.ReportAllHealth); err != nil {
		log.Fatal(err)
	}
	c.Start()
	select {}
}
