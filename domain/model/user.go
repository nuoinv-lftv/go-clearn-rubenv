package model

import "time"

// User Define the structure of the User Model
// type User struct {
// 	ID        int
// 	Name      string
// 	Age       string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt time.Time
// }
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Age       string    `json:"age"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:null" json:"deleted_at"`
}

func (User) TableName() string { return "users" }
