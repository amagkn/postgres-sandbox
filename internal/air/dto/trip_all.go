package dto

import "time"

type TripAllInput struct {
	Town string `json:"town" validate:"min=3"`
}

type TripAllOutput struct {
	ID        int       `json:"id"`
	CompanyID int       `json:"companyID"`
	Plane     string    `json:"plane"`
	TownFrom  string    `json:"townFrom"`
	TownTo    string    `json:"townTo"`
	TimeOut   time.Time `json:"timeOut"`
	TimeIn    time.Time `json:"timeIn"`
}
