package models

type Instructor struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Bio          string `json:"bio"`
	Specialty    string `json:"specialty"`
	Availability string `json:"availability"`
}
