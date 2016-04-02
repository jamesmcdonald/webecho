/*
 * webtest
 * James McDonald <james@jamesmcdonald.com>
 *
 * Periodically hit a URL and log the response time.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestInfo struct {
	url       string
	status    int
	timestamp time.Time
	duration  time.Duration
}

type PollConfig struct {
	url       string
	frequency int
}

func TimeURL(url string) (int, time.Duration) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return 999, 0
	}
	return response.StatusCode, time.Since(start)
}

func Fetcher(url string, frequency int, c chan<- RequestInfo) {
	for {
		status, duration := TimeURL(url)
		c <- RequestInfo{url, status, time.Now(), duration}
		time.Sleep(time.Duration(frequency) * time.Second)
	}
}

func Logger(c <-chan RequestInfo) {
	hostname, _ := os.Hostname()
	for {
		ri := <-c
		log.Printf("%s %s %d %v", hostname, ri.url, ri.status, ri.duration)
	}
}

func ParseConfig(config string) ([]PollConfig, error) {
	var urls []PollConfig

	for _, u := range strings.Split(config, "::") {
		parts := strings.Split(u, ";")
		if len(parts) != 2 {
			return nil, fmt.Errorf("ParseConfig: config %q does not match <url>;<frequency> (eg http://shee.sh/;300)", u)
		}
		freq, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("ParseConfig: frequency %q is not an integer", parts[1])
		}
		p := PollConfig{parts[0], freq}
		urls = append(urls, p)
	}

	return urls, nil
}

func main() {
	urlconfig := os.Getenv("URLS")
	if urlconfig == "" {
		log.Fatal("You need to set the URLS environment variable")
	}

	urls, err := ParseConfig(urlconfig)
	log.Print("Starting monitoring for %v", urls)
	if err != nil {
		log.Fatalf("Could not parse configuration %q: %s", urlconfig, err)
	}

	c := make(chan RequestInfo)

	for _, url := range urls {
		go Fetcher(url.url, url.frequency, c)
	}

	Logger(c)
}
