package models

import (
	"time"
)

type ProjectMetadata struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
