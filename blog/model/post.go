package model

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:100" json:"title"`
	Content   string    `gorm:"size:1000" json:"content"`
	AuthorID  uint      `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
	Status    string    `gorm:"size:20" json:"status"` // 文章状态
}
