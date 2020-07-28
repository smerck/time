package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
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

	go func() {
		for {
			select {
			case <-done:
				return
			case start := <-ticker.C:
				go func() {
					resp, err := c.Get(url)
					if err != nil {
						log.Fatalf("Server Error: %s", err)
					}
					defer resp.Body.Close()
					respChan <- resp
				}()
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
	done <- true
	log.Println("Test complete")
	sr := float64((count-failed)/count) * 100
	log.Printf("Overall Success Rate: %.2f%% (%d/%d)", sr, (count - failed), count)
	log.Printf("Total Number of Requests: %d", count)
}

func parseFlags(rate *int64, url *string, duration *int) {
	flag.Int64Var(rate, "rps", 5, "Number of requests that can be sent per second.")
	flag.StringVar(url, "host", "http://localhost:9001", "URL (Note:including protocol - http[s]://)")
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
