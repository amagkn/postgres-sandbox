package entity

import "time"

type Trip struct {
	ID        int
	CompanyID int
	Plane     string
	TownFrom  string
	TownTo    string
	TimeOut   time.Time
	TimeIn    time.Time
}
