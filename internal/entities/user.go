package entities

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `json:"id"`
	Handler    string    `json:"handler"`
	Email      string    `json:"email"`
	Bio        string    `json:"bio"`
	ProfileUrl string    `json:"profile_url"`
}
