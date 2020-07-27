package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	var rate float64
	var url string
	var testDuration time.Duration
	var count, failed int
	parseFlags(&rate, &url, &testDuration)

	tickRate := time.Duration(1/rate*1000) * time.Millisecond
	ticker := time.NewTicker(tickRate)
	done := make(chan bool)
	c := http.Client{Timeout: time.Second * 2}

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				start := time.Now()
				resp, err := c.Get(url)
				if err != nil {
					log.Fatalf("Server Error: %s", err)
				}
				defer resp.Body.Close()
				d := time.Since(start)
				status := "succeeded"
				if resp.StatusCode != 200 {
					status = "failed"
					failed++
				}
				count++
				log.Printf("Request %s with Status Code %d after %v\n", status, resp.StatusCode, d)

			}
		}
	}()

	time.Sleep(testDuration)
	ticker.Stop()
	done <- true
	fmt.Println("Test complete")
	sr := float64((count-failed)/count) * 100
	log.Printf("Overall Success Rate: %.2f%%", sr)
}

func parseFlags(rate *float64, url *string, duration *time.Duration) {
	flag.Float64Var(rate, "rps", 5, "number of requests that can be sent per second.")
	flag.StringVar(url, "host", "http://localhost:9001", "hostname, including http[s]://")
	flag.DurationVar(duration, "duration", time.Duration(120)*time.Second, "duration of test")
	flag.Parse()
	if !strings.HasPrefix(*url, "http://") {
		log.Fatal("Specify protocol in hostname")
	}
	if *rate > 200 || *rate < 1 {
		log.Fatal("Rate must be between 1 and 100")
	}
	log.Printf("Starting test: Sending %.0f rps to %s for %v", *rate, *url, duration)
}
