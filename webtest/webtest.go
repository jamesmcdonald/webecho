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
	"time"
)

func TimeURL(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Fetcher(url string, frequency uint32, c chan<- time.Duration) {
	for {
		c <- TimeURL(url)
		time.Sleep(time.Duration(frequency) * time.Second)
	}
}

func main() {
	hostname, _ := os.Hostname()
	c := make(chan time.Duration)
	go Fetcher("http://shee.sh/", 10, c)
	for {
		log.Printf("%s %v", hostname, <-c)
	}
}
