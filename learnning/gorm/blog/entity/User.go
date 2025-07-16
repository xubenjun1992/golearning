package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:name`
	Age         int    `json:age`
	Sex         string `json:sex`
	PostList    []Post `gorm:"forenginKey:UserId"`
	UserPostNum int    `json:"userPostNum"`
}
