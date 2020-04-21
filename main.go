package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Rodrigolpb/GoMonitor/models"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan string)

	for _, m := range monitors {
		wg.Add(1)
		go func(m models.Monitor, wg *sync.WaitGroup, ch chan<- string) {
			fmt.Println(m)
			m.Start(ch)
			wg.Done()
		}(m, wg, ch)
	}

	go func(ch <-chan string) {
		for logString := range ch {
			//TODO: Send e-mail with info
			fmt.Printf("Found Problems\n %s", logString)
		}
	}(ch)
	wg.Wait()
}

var monitors = []models.Monitor{
	{
		URL:            "https://random-status-code.herokuapp.com",
		IntervalTime:   time.Duration(1 * time.Minute),
		MaxResposeTime: time.Duration(2 * time.Second),
	},
	{
		URL:            "https://golang.org/",
		IntervalTime:   time.Duration(1 * time.Minute),
		MaxResposeTime: time.Duration(2 * time.Second),
	},
	{
		URL:            "http://google.com/",
		IntervalTime:   time.Duration(1 * time.Minute),
		MaxResposeTime: time.Duration(2 * time.Second),
	},
	{
		URL:            "https://github.com/",
		IntervalTime:   time.Duration(1 * time.Minute),
		MaxResposeTime: time.Duration(2 * time.Second),
	},
	{
		URL:            "https://www.wuxiaworld.com/",
		IntervalTime:   time.Duration(1 * time.Minute),
		MaxResposeTime: time.Duration(2 * time.Second),
	},
}
