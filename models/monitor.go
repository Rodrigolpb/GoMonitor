package models

import (
	"net/http"
	"time"
)

//Monitor - defines a new URL to be monitored and the interval
type Monitor struct {
	URL             string
	IntervalMinutes time.Duration
}

//Start - Method to start sending requests to the monitor's URL
func (m *Monitor) Start() {
	for {
		logRegister := m.testURL()
		logRegister.save()
		time.Sleep(m.IntervalMinutes)
	}
}

func (m *Monitor) testURL() *logRegister {
	startRequest := time.Now()
	resp, _ := http.Get(m.URL)
	logRegister := logRegister{
		url:          m.URL,
		date:         time.Now(),
		statusCode:   *&resp.StatusCode,
		responseTime: time.Since(startRequest),
	}
	return &logRegister
}
