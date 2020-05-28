package models

//Wod is a model
type Wod struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Workout string `json:"workout"`
}
