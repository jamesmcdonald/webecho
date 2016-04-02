/*
 * webtest
 * James McDonald <james@jamesmcdonald.com>
 *
 * Periodically hit a URL and log the response time.
 */

package main

import (
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
	response, _ := http.Get(url)
	return response.StatusCode, time.Since(start)
}

func Fetcher(url string, frequency int, c chan<- RequestInfo) {
	for {
		status, duration := TimeURL(url)
		c <- RequestInfo{url, status, time.Now(), duration}
		time.Sleep(time.Duration(frequency) * time.Second)
	}
}

func main() {

	var urls []PollConfig
	urlconfig := os.Getenv("URLS")
	if urlconfig == "" {
		log.Fatal("You need to set the URLS environment variable")
	}
	for _, u := range strings.Split(urlconfig, ",") {
		parts := strings.Split(u, ";")
		if len(parts) != 2 {
			log.Fatal("Each URL must be in the format url;frequency. Eg http://shee.sh/;300")
		}
		freq, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("The frequency has to be an integer number of seconds")
		}
		p := PollConfig{parts[0], freq}
		urls = append(urls, p)
	}

	hostname, _ := os.Hostname()
	c := make(chan RequestInfo)
	for _, url := range urls {
		go Fetcher(url.url, url.frequency, c)
	}
	for {
		ri := <-c
		log.Printf("%s %s %d %v", hostname, ri.url, ri.status, ri.duration)
	}
}
