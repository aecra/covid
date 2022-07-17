package main

import (
	"github.com/robfig/cron"
)

func corn() {
	c := cron.New()
	c.AddFunc("57 10 8,12 * * * *", report)
	c.Start()
}
