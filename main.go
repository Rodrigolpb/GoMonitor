package main

import (
	"fmt"

	"github.com/Rodrigolpb/GoMonitor/models"
)

func main() {
	for _, m := range monitors {
		fmt.Println(m)
		m.Start()
	}
}

var monitors = []models.Monitor{
	{
		ID:              1,
		URL:             "https://random-status-code.herokuapp.com",
		IntervalMinutes: 1,
	},
	{
		ID:              2,
		URL:             "http://qualiex.com.br/",
		IntervalMinutes: 5,
	},
	{
		ID:              3,
		URL:             "https://www.alura.com.br/",
		IntervalMinutes: 5,
	},
}
