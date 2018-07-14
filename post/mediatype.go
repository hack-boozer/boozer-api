package post

import (
	"time"

	"github.com/hack-boozer/boozer-api/account"
)

// Media post media
type Media struct {
	User      account.Media `json:"user"`
	Comment   string        `json:"comment"`
	Rate      float64       `json:"rating"`
	Photo     string        `json:"photo"`
	CreatedAt time.Time     `json:"createdAt"`
}
