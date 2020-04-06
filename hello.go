package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	monitorings := 10
	interval := 2

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
			config := map[string]int{
				"Number of Monitorings": monitorings,
				"Interval (Minutes)":    interval,
			}
			printConfiguration(config)
		case 4:
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
	fmt.Println("3 - Show Configurations")
	fmt.Println("4 - Exit")
	fmt.Println("@Version ", version)
	fmt.Print("Chose option: ")
	fmt.Scan(&cmdOption)
	fmt.Println("-----------------------")

	return cmdOption
}

func startMonitoring(monitorings int, interval int) {
	sites := getSitesFromFile()

	fmt.Println("Started Monitoring...")
	for i := 0; i < monitorings; i++ {
		fmt.Println("Monitoring N:", i+1)
		for i, site := range sites {
			fmt.Printf("> Test Sites[%d] - ", i)
			testSite(site)
		}
		time.Sleep(time.Duration(interval) * time.Minute)
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
	filePath := "logs/log" + time.Now().Format("02-01-2006") + ".txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		exitProgram(true)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Running: " + strconv.FormatBool(isRunning) + "\n")
	file.Close()
}

func showLogs() {
	filePath := "logs/log" + time.Now().Format("02-01-2006") + ".txt"
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		exitProgram(false)
	}

	fmt.Println("Logs from file:", filePath)
	fmt.Println()
	fmt.Println(string(file))
}

func exitProgram(hasError bool) {
	fmt.Println("Exiting...")
	if hasError {
		os.Exit(-1)
	}
	os.Exit(0)
}

func printConfiguration(configurations map[string]int) {
	for i, config := range configurations {
		fmt.Println(i, ": ", config)
	}
}
