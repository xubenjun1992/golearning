package main

import (
	"blog/controller"
	"blog/core"
	"blog/dao"
	"blog/model"
	"blog/rout"
	"blog/service"
)

func main() {
	db := core.GetDb()
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)

	userDao := dao.NewUserDao()
	commentDao := dao.NewCommentDao()
	articleDao := dao.NewArticleDao()
	userService := service.NewUserService(userDao)
	commentService := service.NewCommentService(commentDao)
	articleService := service.NewArticleService(articleDao)

	r := rout.SetBlogRouts(
		controller.NewArticleController(articleService),
		controller.NewCommentController(commentService),
		controller.NewUserController(userService),
	)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
