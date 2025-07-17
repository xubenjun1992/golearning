package reqdto

type CreateCommentDTO struct {
	Id        uint   `json:"id"`
	Content   string `json:"content" binding:"required"`
	ArticleId uint   `json:"article_id" binding:"required"`
}
