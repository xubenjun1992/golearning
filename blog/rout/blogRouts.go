package rout

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetBlogRouts(articleController controller.ArticleController, commentController controller.CommentController, userController controller.UserController) *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	{
		public.POST("/blog/login", userController.Login)
		public.POST("/blog/register", userController.RegistUser)
		public.GET("/blog/articles", articleController.GetAllArticles)
		public.GET("/blog/articles/:id", articleController.GetArticleById)
		public.GET("/blog/comments/article/:articleId", commentController.GetCommentsByArticleId)
	}

	private := r.Group("/api")
	private.Use(middleware.AuthMiddleware())
	{
		private.POST("/blog/articles", articleController.CreateArticle)
		private.PUT("/blog/articles/:id", articleController.UpdateArticle)
		private.DELETE("/blog/articles/:id", articleController.DeleteArticle)
		private.GET("/blog/articles/user/:userId", articleController.GetArticlesByUserId)
		private.POST("/blog/comments", commentController.CreateComment)
		private.GET("/blog/comments/:id", commentController.GetCommentById)
		private.PUT("/blog/comments/:id", commentController.UpdateComment)
		private.DELETE("/blog/comments/:id", commentController.DeleteComment)
	}
	return r
}
