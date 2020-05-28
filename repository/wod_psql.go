package wodrepository

import (
	"wod-api/driver"
	"wod-api/models"
)

// WodRepository is a struct
type WodRepository struct{}

var db = driver.ConnectDB()

//Store is to save a wod in database
func (w WodRepository) Store(wod *models.Wod) (*models.Wod, error) {
	stmt := "insert into wod (title, workout) values($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, wod.Title, wod.Workout).Scan(&wod.ID)

	if err != nil {
		return wod, err
	}

	return wod, err
}

//Get is for retrieve a wod from db
func (w WodRepository) Get() (models.Wod, error) {
	var wod models.Wod
	row := db.QueryRow("SELECT title,workout FROM wod ORDER BY RANDOM() LIMIT 1")
	err := row.Scan(&wod.Title, &wod.Workout)

	if err != nil {
		return wod, err
	}

	return wod, nil
}
