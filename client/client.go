package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	var rate int64
	var url string
	var duration, count, failed int
	parseFlags(&rate, &url, &duration)

	tickRate := time.Duration(1/float64(rate)*1000) * time.Millisecond
	ticker := time.NewTicker(tickRate)
	done := make(chan bool)
	c := http.Client{Timeout: time.Second * 2}
	respChan := make(chan *http.Response)
	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case <-done:
				return
			case start := <-ticker.C:
				go sendRequest(&c, respChan, &wg, url)
				resp := <-respChan
				d := time.Since(start)
				status := "succeeded"
				if resp.StatusCode != 200 {
					status = "failed"
					failed++
				}
				count++
				log.Printf("Request %d %s with Status Code %d after %v\n", count, status, resp.StatusCode, d)
			}
		}
	}()

	time.Sleep(time.Duration(duration) * time.Second)
	ticker.Stop()
	wg.Wait()
	done <- true
	log.Println("Test complete")
	sr := float64((count-failed)/count) * 100
	log.Printf("Overall Success Rate: %.2f%% (%d/%d)", sr, (count - failed), count)
	log.Printf("Total Number of Requests: %d", count)
}

func parseFlags(rate *int64, url *string, duration *int) {
	flag.Int64Var(rate, "rps", 5, "Number of requests that client will send per second")
	flag.StringVar(url, "host", "http://localhost:9001", "Server URL (Note: including protocol)")
	flag.IntVar(duration, "duration", 120, "Duration of test")
	flag.Parse()

	if !strings.HasPrefix(*url, "http://") {
		log.Fatal("Specify protocol in hostname.")
	}
	if *rate > 1000 || *rate < 1 {
		log.Fatal("Rate must be between 1 and 1000.")
	}
	if *duration < 0 {
		log.Fatal("Test duration must be greater than 0.")
	}
	log.Printf("Starting test: Sending %d requests per second to %s for %ds.", *rate, *url, *duration)
}

func sendRequest(c *http.Client, ch chan *http.Response, wg *sync.WaitGroup, url string) {
	wg.Add(1)
	resp, err := c.Get(url)
	defer wg.Done()
	if err != nil {
		log.Printf("client error: %s", err)
	}
	defer resp.Body.Close()
	ch <- resp
}
