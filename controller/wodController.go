package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"wod-api/models"
	wodrepository "wod-api/repository"
)

// NewWodFromRequest creates new Wod from http request
func NewWodFromRequest(r *http.Request) (*models.Wod, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var wod models.Wod

	err = json.Unmarshal(body, &wod)
	if err != nil {
		return nil, err
	}

	if wod.Title == "" || wod.Workout == "" {
		return nil, errors.New("title or workout cannot be empty")
	}

	return &wod, nil
}

// StoreWodInDb stores a new wod in db
func StoreWodInDb(wod *models.Wod) error {
	wodRepo := wodrepository.WodRepository{}
	wod, err := wodRepo.Store(wod)
	return err
}

// GetRamdomWodFromDb retrieves a ramdom workout for db
func GetRamdomWodFromDb() (*models.Wod, error) {
	wodRepo := wodrepository.WodRepository{}
	wod, err := wodRepo.Get()
	if err != nil {
		return nil, err
	}
	return &wod, nil
}
