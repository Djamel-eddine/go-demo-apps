package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
) 

var port = "5000"

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world, %q \n", html.EscapeString(r.URL.Path))
    })

    startServer()
}

func startServer()  {
    
    log.Printf("Listening on port %s", port)
    err := http.ListenAndServe(":"+port, nil)
    
    if err != nil {
        log.Fatal(err)
    }
}