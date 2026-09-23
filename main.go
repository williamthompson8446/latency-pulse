package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type ProbeResult struct {
	URL      string
	Duration time.Duration
	Status   int
	Err      error
}

func probeURL(url string, client *http.Client, ch chan<- ProbeResult, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	resp, err := client.Get(url)
	dur := time.Since(start)

	if err != nil {
		ch <- ProbeResult{URL: url, Duration: dur, Status: 0, Err: err}
		return
	}
	defer resp.Body.Close()
	ch <- ProbeResult{URL: url, Duration: dur, Status: resp.StatusCode, Err: nil}
}

func main() {
	urls := []string{
		"https://httpbin.org/status/200",
		"https://api.github.com",
		"https://cloudflare.com",
	}

	if len(os.Args) > 1 {
		urls = os.Args[1:]
	}

	client := &http.Client{Timeout: 6 * time.Second}
	ch := make(chan ProbeResult, len(urls))
	var wg sync.WaitGroup

	fmt.Println("==================================================")
	fmt.Printf("  Goroutine HTTP Probe — Concurrent Latency Checker
")
	fmt.Println("==================================================")

	for _, u := range urls {
		wg.Add(1)
		go probeURL(u, client, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for res := range ch {
		if res.Err != nil {
			fmt.Printf("[-] %-32s -> ERR: %v
", res.URL, res.Err)
		} else {
			fmt.Printf("[+] %-32s -> %d OK (took %v)
", res.URL, res.Status, res.Duration.Round(time.Millisecond))
		}
	}
	fmt.Println("==================================================")
}
