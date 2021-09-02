package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	j := cron.New(cron.WithSeconds(),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
		cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)))

	_, err := j.AddFunc("@every 10s", func() {
		fmt.Println("Start...")
	})
	if err != nil {
		panic(err)
	}

	j.Run()
}
