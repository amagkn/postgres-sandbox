package dto

type PassengerAllInput struct {
	Suffix string `json:"suffix" validate:"omitempty,min=2"`
}

type PassengerAllOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
