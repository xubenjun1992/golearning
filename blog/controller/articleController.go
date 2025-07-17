package controller

import (
	"blog/enum"
	"blog/model"
	"blog/reqdto"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	CreateArticle(c *gin.Context)
	GetArticleById(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
	GetArticlesByUserId(c *gin.Context)
	GetAllArticles(c *gin.Context)
}

type articleController struct {
	articleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) ArticleController {
	return &articleController{
		articleService: articleService,
	}
}

func (a *articleController) CreateArticle(c *gin.Context) {
	var createArticleDTO reqdto.CreateArticleDTO
	c.ShouldBindJSON(&createArticleDTO)
	var userId uint
	if createArticleDTO.UserId == 0 {
		userId = c.GetUint("userId")
	} else {
		userId = createArticleDTO.UserId
	}
	post := model.Post{
		Title:    createArticleDTO.Title,
		Content:  createArticleDTO.Content,
		AuthorID: userId,
		Status:   enum.ArticleStatusDraft,
	}
	if post, err := a.articleService.CreateArticle(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Article created successfully", "article": post})
	}
}

func (a *articleController) GetArticleById(c *gin.Context) {
	articleIdStr := c.Param("id")
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	if post, err := a.articleService.GetArticleById(uint(articleId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"article": post})
	}
}
func (a *articleController) UpdateArticle(c *gin.Context) {
	var createArticleDTO reqdto.CreateArticleDTO
	c.ShouldBindJSON(&createArticleDTO)

	post, err := a.articleService.GetArticleById(createArticleDTO.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if enum.ArticleStatusDraft != post.Status {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only draft articles can be updated"})
		return
	}
	post.Title = createArticleDTO.Title
	post.Content = createArticleDTO.Content
	if post, err := a.articleService.UpdateArticle(*post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully", "article": post})
	}
}
func (a *articleController) DeleteArticle(c *gin.Context) {
	articleIdStr := c.Param("id")
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}
	if err := a.articleService.DeleteArticle(uint(articleId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
	}
}
func (a *articleController) GetArticlesByUserId(c *gin.Context) {
	uid, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}
	userId, ok := uid.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	if posts, err := a.articleService.GetArticlesByUserId(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"articles": posts})
	}
}
func (a *articleController) GetAllArticles(c *gin.Context) {
	if posts, err := a.articleService.GetAllArticles(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"articles": posts})
	}
}
