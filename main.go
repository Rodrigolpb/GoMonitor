package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Rodrigolpb/GoMonitor/models"
)

func main() {
	wg := &sync.WaitGroup{}
	for _, m := range monitors {
		wg.Add(1)
		go func(m models.Monitor, wg *sync.WaitGroup) {
			fmt.Println(m)
			m.Start()
		}(m, wg)
	}
	wg.Wait()
}

var monitors = []models.Monitor{
	{
		URL:             "https://random-status-code.herokuapp.com",
		IntervalMinutes: time.Duration(1 * time.Minute),
	},
	{
		URL:             "https://golang.org/",
		IntervalMinutes: time.Duration(1 * time.Minute),
	},
	{
		URL:             "http://google.com/",
		IntervalMinutes: time.Duration(1 * time.Minute),
	},
	{
		URL:             "https://github.com/",
		IntervalMinutes: time.Duration(1 * time.Minute),
	},
	{
		URL:             "https://www.wuxiaworld.com/",
		IntervalMinutes: time.Duration(1 * time.Minute),
	},
}
