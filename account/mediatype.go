package account

import uuid "github.com/satori/go.uuid"

// Media account media
type Media struct {
	ID       uuid.UUID `json:"id"`
	NickName string    `json:"nickName"`
	Email    string    `json:"email"`
	Photo    string    `json:"photo"`
}
