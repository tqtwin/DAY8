// package models/user.go
package models

import "time"

// User là cấu trúc đại diện cho người dùng
type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"password"`          // Thêm trường Password
	Bio       string     `json:"bio"`               // Thêm trường Bio
	Posts     []Post     `gorm:"foreignKey:UserID"` // Thiết lập quan hệ với Post
}
