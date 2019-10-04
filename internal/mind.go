package model

import "context"

// Mind represents the mind of users
type Mind struct {
	Base
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsDelete bool   `json:"is_delete"`

	User   *User `json:"user"`
	UserID int   `json:"user_id"`
}

// MindDB represents the mind database interface
type MindDB interface {
	View(context.Context, int) (*User, error)
}
