package models

import "time"

type Asset struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
