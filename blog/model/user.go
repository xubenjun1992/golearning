package model

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:100" json:"username"`
	Email     string    `gorm:"size:100" json:"email"`
	Password  string    `gorm:"size:100" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Posts     []Post    `gorm:"foreignKey:AuthorID" json:"posts"`
	Comments  []Comment `gorm:"foreignKey:AuthorID" json:"comments"`
}
