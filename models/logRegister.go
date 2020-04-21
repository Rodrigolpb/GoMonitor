package models

import (
	"fmt"
	"os"
	"time"
)

//logRegister - defines a new log entry
type logRegister struct {
	url          string
	date         time.Time
	statusCode   int
	responseTime time.Duration
}

const basicFilePath = "logs/log_"

//save - save log content into text file
func (l *logRegister) save() {
	filePath := basicFilePath + time.Now().Format("02-01-2006") + ".txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	file.WriteString(l.String())
}

func (l *logRegister) String() string {
	return fmt.Sprintf("URL: %v - Status Code: %v - Date: %v - Response Time: %vs\n", l.url, l.statusCode, l.date.Format("02/01/2006 15:04:05"), l.responseTime.Seconds())
}
