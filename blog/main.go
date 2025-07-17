package main

import (
	"blog/controller"
	"blog/dao"
	"blog/rout"
	"blog/service"
)

func main() {
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
