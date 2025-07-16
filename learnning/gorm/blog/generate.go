package blog

import (
	"encoding/json"
	"errors"
	"log"

	"gorm.io/gorm"
	"main.go/gorm/blog/entity"
	"main.go/gorm/samplesql/core"
)

func GenerateTable() {
	err := core.DB.AutoMigrate(&entity.User{}, &entity.Post{}, &entity.Comment{})
	if err != nil {
		panic("创建表错误")
	}
	log.Printf("创建表成功")
}

func QueryPostsAndCommentsByUserId(userId uint) {
	user := entity.User{}
	err := core.DB.Preload("PostList.CommentList", func(db *gorm.DB) *gorm.DB {
		return db.Order("create_at desc")
	}).First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("未查找到用户")
			return
		} else {
			log.Panicf("查询用户错误:%v", err)
		}

	}
	userData, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Panicf("序列化User错误，%s", err)
		return
	}
	log.Printf("查询结果:%v", string(userData))
}

func QueryMostCommentsPost() *entity.Post {
	var post entity.Post
	err := core.DB.Model(&entity.Post{}).Table("posts p").Select("p.*, count(1)").Joins("left join comments c on p.id = c.post_id").Group("p.id").Order("count(1) desc").Limit(1).Take(&post).Error
	if err != nil {
		log.Printf("查询错误:%v", err)
		return &post
	}
	data, _ := json.MarshalIndent(&post, "", "  ")
	log.Printf("评论数最多的文章是：%v", string(data))
	return &post
}
