package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func task() {
	fmt.Println("Task is being performed.")
}

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(task) // every 5 seconds

	// cron expressions supported
	s.Cron("*/1 * * * *").Do(task) // every minute

	// starts the scheduler asynchronously
	s.StartAsync()
}
