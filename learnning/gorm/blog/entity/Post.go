package entity

import (
	"log"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string    `json:title`
	Content     string    `json:content`
	CommentList []Comment `gorm:"forenginKey:PostId" json:"commentList"`
	UserId      uint      `json:userId`
	Status      string    `json:status`
}

func (p *Post) AfterCreate(db *gorm.DB) {
	log.Printf("创建后执行钩子函数")
	db.Model(&User{}).UpdateColumn("user_post_num", gorm.Expr("user_post_num = user_post_num + 1")).Where("user_id = ?", p.UserId)
	log.Panicf("更新用户文章数钩子执行完毕")
}
