package photo

import uuid "github.com/satori/go.uuid"

// Photo photo
type Photo struct {
	PostID uuid.UUID `json:"postId" sql:"type:uuid"`
	File   string    `json:"file"`
}
