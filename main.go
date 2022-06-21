package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
)

const APPNAME string = "whereru"
const ICON string = "assets/icon.png"

func main() {
	stopHour := flag.Int("stop", 0, "hour at which the program should stop")

	beeep.Notify(APPNAME, "Session starting :)", ICON)

	s := gocron.NewScheduler(time.UTC)

	s.Every(23).Minutes().WaitForSchedule().Do(func() {
		if coinToss() {
			alert()
		}
	})

	s.StartBlocking()
	maybeStop(*stopHour, s)
}

func coinToss() bool {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64()
	return random < 0.23
}

func alert() {
	err := beeep.Alert(APPNAME, "Take a minute to become aware of your body, emotions, and thoughts.", ICON)
	if err != nil {
		panic(err)
	}
}

func maybeStop(hour int, s *gocron.Scheduler) {
	if hour != 0 {
		cronFmtTime := fmt.Sprintf("0 %d * * *", hour)
		s.Cron(cronFmtTime).Do(func() {
			s.Stop()
		})
	}
}
