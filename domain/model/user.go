package model

import "time"

// User Define the structure of the User Model
type User struct {
	ID        uint      `gorm:"primary_key" join:"id"`
	Name      string    `join:"name"`
	Age       string    `join:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
