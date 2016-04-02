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
	"time"
)

func TimeURL(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func main() {
	log.Print(TimeURL("http://shee.sh/"))
}
