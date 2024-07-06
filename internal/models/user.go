package models

import "time"

type User struct {
	ID              int       `json:"id" db:"id"`
	Username        string    `json:"username" db:"username"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	DefaultLocation string    `json:"default_location" db:"default_location"`
	PreferredUnit   string    `json:"preferred_unit" db:"preferred_unit"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
