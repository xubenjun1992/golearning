package entity

import (
	"log"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	CommentInfo string `json:commentInfo`
	PostId      uint   `json:postId`
}

func (comment *Comment) AfterDelete(db *gorm.DB) {
	log.Printf("执行删除后钩子")
	var count int64
	err := db.Model(&Comment{}).Where("post_id = ?", comment.PostId).Count(&count).Error
	if err != nil {
		log.Panicf("钩子函数查询文章评论数失败:%v", err)
		return
	}

	if count == 0 {
		err := db.Model(Post{}).Where("id = ?", comment.PostId).UpdateColumn("status", "无评论").Error
		if err != nil {
			log.Panicf("更新文章评论状态失败,%v", err)
			return
		}
		log.Printf("评论数为0,已更新文章 %d 状态为“已发布”", comment.PostId)
	}
}
