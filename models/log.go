package models

import (
	"fmt"
	"os"
	"time"
)

//Log defines a new log to be saved
type Log struct {
	url          string
	date         time.Time
	statusCode   int
	err          error
	responseTime time.Duration
}

const basicFilePath = "logs/log"

//Save - saves log content into log file
func (l *Log) Save() {
	filePath := basicFilePath + time.Now().Format("02-01-2006") + ".txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	file.WriteString(l.string())
}

func (l *Log) string() string {
	return fmt.Sprintf("URL: %v - Status Code: %v - Date: %v - Response Time: %vs - Error: %v\n", l.url, l.statusCode, l.date.Format("02/01/2006 15:04:05"), l.responseTime.Seconds(), l.err)
}
