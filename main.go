package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "5151"
)

func main() {
	var (
		port string
	)

	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
	log.SetOutput(os.Stdout)
	log.Printf("Starting on port %s", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		body := fmt.Sprintf("%s: %s", k, v)
		body = fmt.Sprintln(body)
		bodyBytes := []byte(body)
		w.Write(bodyBytes)
	}
}
