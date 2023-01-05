package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func taskEvery5Seconds() {
	log.Println("Task is every 5 seconds")
}

func taskEveryMinute() {
	log.Println("Task is every minute")
}

func main() {
	s := gocron.NewScheduler(time.UTC)

	// Cron expressions supported: https://crontab.guru/
	s.Cron("*/1 * * * *").Do(taskEveryMinute) // every minute

	// every 24th hour
	s.Cron("0 */24 * * *").Do(func() {
		log.Println("Task is every 24 hours")
	})

	// starts the scheduler blocking
	s.StartBlocking()
}
