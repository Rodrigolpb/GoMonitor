package models

import (
	"net/http"
	"time"
)

//Monitor defines a new URL to be monitored and the interval
type Monitor struct {
	ID              int32
	URL             string
	IntervalMinutes int32
}

//Start - Method to start sending requests to the monitor's URL
func (m *Monitor) Start() {
	for {
		log := m.testURL()
		log.Save()
		time.Sleep(time.Duration(m.IntervalMinutes) * time.Minute)
	}
}

func (m *Monitor) testURL() *Log {
	startRequest := time.Now()
	resp, err := http.Get(m.URL)
	log := Log{
		url:          m.URL,
		date:         time.Now(),
		statusCode:   *&resp.StatusCode,
		err:          err,
		responseTime: time.Since(startRequest),
	}
	return &log
}
