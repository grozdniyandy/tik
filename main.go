package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var checkedLines int
var mu sync.Mutex

func main() {
	customTextDone := make(chan bool)

	go displayCustomText(customTextDone)

	<-customTextDone

	fmt.Print("\n\nEnter the path to the file containing domains: ")
	var inputFileName string
	_, err := fmt.Scanln(&inputFileName)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	fmt.Print("Enter the number of concurrent threads: ")
	var numThreads int
	_, err = fmt.Scanln(&numThreads)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup

	concurrencyLimit := make(chan struct{}, numThreads)

	totalLines := 0
	for scanner.Scan() {
		totalLines++
	}
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	progressFormat := "Progress: [%s] %d/%d"
	loadingChars := []string{"|", "/", "-", "\\"}

	fmt.Printf(progressFormat, " ", checkedLines, totalLines)

	for scanner.Scan() {
		domain := scanner.Text()
		if domain != "" {
			wg.Add(1)
			concurrencyLimit <- struct{}{}
			go func(domain string) {
				defer func() {
					<-concurrencyLimit
					wg.Done()
					mu.Lock()
					checkedLines++
					progress := loadingChars[checkedLines%len(loadingChars)]
					fmt.Printf(progressFormat+"\r", progress, checkedLines, totalLines)
					mu.Unlock()
				}()
				checkSecurityHeaders(domain)
			}(domain)
		}
	}

	wg.Wait()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
	}
	fmt.Println("\nDone!")
}

func displayCustomText(customTextDone chan<- bool) {
	customText := "Made with  ˗ˋˏ ♡ ˎˊ˗ by GrozdniyAndy of XSS.is"

	for i := 0; i <= len(customText); i++ {
		fmt.Print("\r" + customText[:i] + "_")
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print("\r" + customText)
	time.Sleep(1000 * time.Millisecond)

	customTextDone <- true 
}


func checkSecurityHeaders(domain string) {
	checkURL(domain, "http")
	checkURL(domain, "https")
}

func checkURL(domain, protocol string) {
	url := protocol + "://" + domain
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	contentSecurityPolicy := resp.Header.Get("Content-Security-Policy")
	xFrameOptions := resp.Header.Get("X-Frame-Options")

	if contentSecurityPolicy == "" && xFrameOptions == "" {
		fmt.Printf("Success: %s - %s - Headers not found (Content-Security-Policy and X-Frame-Options)\n", domain, protocol)
	}
}
