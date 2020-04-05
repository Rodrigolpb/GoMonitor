package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	const monitorings = 5
	const interval = 5
	version := "1.0"
	getSitesFromFile()

	for {
		cmdOption := readMenuOption(version)
		switch cmdOption {
		case 1:
			startMonitoring(monitorings, interval)
		case 2:
			showLogs()
		case 3:
			exitProgram(false)
		default:
			fmt.Println("Invalid option. Please, try again...")
			fmt.Println()
		}
	}
}

func readMenuOption(version string) int {
	var cmdOption int

	fmt.Println("-----------------------")
	fmt.Println("Menu - Choose an option")
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("3 - Exit")
	fmt.Println("@Version ", version)
	fmt.Print("Chose option: ")
	fmt.Scan(&cmdOption)
	fmt.Println("-----------------------")

	return cmdOption
}

func startMonitoring(monitorings int, interval float32) {
	sites := getSitesFromFile()

	fmt.Println("Started Monitoring...")
	for i := 0; i < monitorings; i++ {
		fmt.Println("Monitoring N:", i+1)
		for i, site := range sites {
			fmt.Printf("> Test Sites[%d] - ", i)
			testSite(site)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
		exitProgram(true)
	}

	fmt.Printf("URL: (%s) STATUS CODE: %d\n", site, res.StatusCode)
	if res.StatusCode == 200 {
		registerLogs(site, true)
	} else {
		registerLogs(site, false)
	}

}

func getSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
		exitProgram(true)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func registerLogs(site string, isRunning bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		exitProgram(true)
	}

	fmt.Println(file)
}

func showLogs() {
	fmt.Println("TODO: Show Logs...")
	fmt.Println()
}

func exitProgram(hasError bool) {
	fmt.Println("Exiting...")
	if hasError {
		os.Exit(-1)
	}
	os.Exit(0)
}
