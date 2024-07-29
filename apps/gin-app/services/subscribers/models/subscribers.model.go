package models

type Subscriber struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	ID      string `json:"id"`
	Country string `json:"country"`
}
