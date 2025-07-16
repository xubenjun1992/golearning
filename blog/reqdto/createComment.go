package reqdto

type CreateCommentDTO struct {
	Id        uint   `json:"id"`
	Content   string `json:"content" binding:"required"`
	ArticleID uint   `json:"article_id" binding:"required"`
}
