/*
 * echoserver
 * James McDonald <james@jamesmcdonald.com>
 * 
 * Echo the URI path in the request back to the client.
 */

package main

import (
    "fmt"
    "log"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, r.URL.Path)
}

func main() {
    http.HandleFunc("/", Handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
