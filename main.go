package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./models"
	"./response"
)

func home(w http.ResponseWriter, r *http.Request) {
	response.JSONSuccess(w, "Welcome to work out of the day API homepage")
}

func wod(w http.ResponseWriter, r *http.Request) {
	var error models.Error
	if r.Method == http.MethodPost {
		newWod(w, r)
	} else if r.Method == http.MethodGet {
		response.JSONSuccess(w, "Do 1000 burpees")
		return
	} else {
		error.Message = "Invalid request method"
		response.JSONError(w, http.StatusMethodNotAllowed, error)
		return
	}
}

func newWod(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var error models.Error
	if err != nil {
		error.Message = "cannot parse request body"
		response.JSONError(w, http.StatusBadRequest, error)
		return
	}

	var wod models.Wod

	err = json.Unmarshal(body, &wod)
	if err != nil {
		error.Message = "cannot unmarshal request body."
		response.JSONError(w, http.StatusBadRequest, error)
		return
	}

	if wod.Title == "" || wod.Workout == "" {
		error.Message = "title or workout cannot be empty"
		response.JSONError(w, http.StatusBadRequest, error)
		return
	}

	response.JSONSuccess(w, "Added wod with title: "+wod.Title)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/work-out-of-the-day", wod)
	fmt.Printf("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
