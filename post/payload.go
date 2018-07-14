package post

import uuid "github.com/satori/go.uuid"

// Create post create payload
type Create struct {
	AccountID uuid.UUID `json:"accountId"`
	Comment   string    `json:"comment"`
	Photos    []string  `json:"photos"`
	Rate      float64   `json:"rating"`
}
