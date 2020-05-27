package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to work out of the day API homepage")
	if err != nil {
		log.Panic(err)
	}
}

func wod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newWod(w, r)
	} else {
		_, err := fmt.Fprintf(w, "Invalid request method")
		if err != nil {
			log.Panic(err)
		}
	}
}

func newWod(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "POST request to /wod endpoint!")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/work-out-of-the-day", wod)
	fmt.Printf("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
