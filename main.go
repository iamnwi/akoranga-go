package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
	}

	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello, This is Cloud Computing Assignment 1!\n")
		fmt.Fprintf(w, "Hostname: [%s]\n", hostname)
		fmt.Fprintf(w, "NetID: win205\n")
	})

	http.ListenAndServe(":"+port, nil)
}
