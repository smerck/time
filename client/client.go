package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//parse flags
	rate := 400

	ticker := time.NewTicker(time.Second / rate)
	done := make(chan bool)

	c := http.Client{Timeout: time.Second * 2}

	 respChan := make(chan http.Response)
	// send request according to rate
	go func() {
		for {
			select {
			case <-done:
				return nil
			case <-ticker.C:
				success := 0
				//var latency time.Duration
				//log.Printf("Sending %d requests\n", rate)
				//for i := 0; i < rate; i++ {
				start := time.Now()
				go c.Get("http://localhost:9001")
				if err != nil {
					log.Fatal("Server Error: %s", err)
				}
				defer resp.Body.Close()
				d := time.Since(start)
				if resp.StatusCode == 200 {
					success++
				}
				latency += d
			}
		}
	}()

    go func() {
		select { 
		case <-success:
			success++
		}
		case <-fail:
			
	}
	
	
	ticker.Stop()
	done <- true
	fmt.Println("Test stopped")

}

func fetch(url, ch chan<- http.Response) {

}
