package atenea

import (
	"time"
)

// User represents a user information returned by UserService
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Token    string `json:"token"`

	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"-"`
}
