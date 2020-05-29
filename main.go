package main

import (
	"fmt"
	"net/http"
	wodController "wod-api/controller"
	"wod-api/models"
	"wod-api/response"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/work-out-of-the-day", wod)
	fmt.Printf("Server listening on port 8081")
	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	response.JSONSuccess(w, "Welcome to work out of the day API homepage")
}

func wod(w http.ResponseWriter, r *http.Request) {
	var error models.Error
	if r.Method == http.MethodPost {
		newWod(w, r)
		return
	}
	if r.Method == http.MethodGet {
		randomWod(w, r)
		return
	}

	error.Message = "Invalid request method"
	response.JSONError(w, http.StatusMethodNotAllowed, error)

}

func newWod(w http.ResponseWriter, r *http.Request) {
	var error models.Error
	wod, err := wodController.NewWodFromRequest(r)

	if err != nil {
		error.Message = err.Error()
		response.JSONError(w, http.StatusUnprocessableEntity, error)
		return
	}
	err = wodController.StoreWodInDb(wod)

	if err != nil {
		error.Message = "Server error"
		response.JSONError(w, http.StatusInternalServerError, error)
		return
	}
	response.JSONSuccess(w, "Added wod with title: "+wod.Title)
}

func randomWod(w http.ResponseWriter, r *http.Request) {
	var error models.Error
	wod, err := wodController.GetRamdomWodFromDb()
	if err != nil {
		error.Message = err.Error()
		response.JSONError(w, http.StatusInternalServerError, error)
		return
	}
	response.JSONSuccess(w, wod)
}
