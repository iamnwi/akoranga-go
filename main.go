package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Kia ora from port "+port)
	})
	http.ListenAndServe(":"+port, nil)
}
