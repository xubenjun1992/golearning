package controller

import (
	"blog/model"
	"blog/reqdto"
	"blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(c *gin.Context)
	GetCommentById(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	GetCommentsByArticleId(c *gin.Context)
	GetCommentsByUserId(c *gin.Context)
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &commentController{
		commentService: commentService,
	}
}

func (c *commentController) CreateComment(ctx *gin.Context) {
	var createCommentDTO reqdto.CreateCommentDTO
	if err := ctx.ShouldBindJSON(&createCommentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userId := ctx.GetUint("userId")
	comment := model.Comment{
		Content:  createCommentDTO.Content,
		PostID:   createCommentDTO.ArticleId,
		AuthorID: userId,
	}

	if createdComment, err := c.commentService.CreateComment(comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment created successfully", "comment": createdComment})
	}
}

func (c *commentController) GetCommentById(ctx *gin.Context) {
	commentId := ctx.GetInt64("id")
	if comment, err := c.commentService.GetCommentById(uint(commentId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"comment": comment})
	}
}

func (c *commentController) UpdateComment(ctx *gin.Context) {
	var updateCommentDTO reqdto.CreateCommentDTO
	if err := ctx.ShouldBindJSON(&updateCommentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	userId := ctx.GetUint("userId")

	comment := model.Comment{
		ID:       updateCommentDTO.Id,
		Content:  updateCommentDTO.Content,
		AuthorID: userId,
	}

	if updatedComment, err := c.commentService.UpdateComment(comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully", "comment": updatedComment})
	}
}

func (c *commentController) DeleteComment(ctx *gin.Context) {
	commentId := ctx.GetInt64("id")
	userId := ctx.GetUint("userId")
	if err := c.commentService.DeleteComment(uint(commentId), uint(userId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
	}
}

func (c *commentController) GetCommentsByArticleId(ctx *gin.Context) {
	articleId := ctx.GetInt64("articleId")
	if comments, err := c.commentService.GetCommentsByArticleId(uint(articleId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"comments": comments})
	}
}

func (c *commentController) GetCommentsByUserId(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	if comments, err := c.commentService.GetCommentsByUserId(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"comments": comments})
	}
}
