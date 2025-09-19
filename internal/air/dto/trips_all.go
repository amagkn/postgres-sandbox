package dto

import "time"

type TripsAllInput struct {
	Town string `json:"town" validate:"min=3"`
}

type TripsAllOutput struct {
	ID        int       `json:"id"`
	CompanyID int       `json:"companyID"`
	Plane     string    `json:"plane"`
	TownFrom  string    `json:"townFrom"`
	TownTo    string    `json:"townTo"`
	TimeOut   time.Time `json:"timeOut"`
	TimeIn    time.Time `json:"timeIn"`
}
