package dao

import (
	"blog/core"
	"blog/model"
)

type CommentDao interface {
	CreateComment(comment model.Comment) (*model.Comment, error)
	DeleteComment(id uint, userId uint) error
	GetCommentById(id uint) (*model.Comment, error)
	GetCommentsByArticleId(articleId uint) ([]model.Comment, error)
	GetCommentsByUserId(userId uint) ([]model.Comment, error)
	UpdateComment(comment model.Comment) (*model.Comment, error)
}

type commentDao struct{}

func NewCommentDao() CommentDao {
	return &commentDao{}
}

func (c *commentDao) CreateComment(comment model.Comment) (*model.Comment, error) {
	db := core.GetDb()
	if err := db.Create(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentDao) DeleteComment(id uint, userId uint) error {
	db := core.GetDb()
	if err := db.Where("author_id = ?", userId).Delete(&model.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (c *commentDao) GetCommentById(id uint) (*model.Comment, error) {
	db := core.GetDb()
	var comment model.Comment
	if err := db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentDao) GetCommentsByArticleId(articleId uint) ([]model.Comment, error) {
	db := core.GetDb()
	var comments []model.Comment
	if err := db.Where("article_id = ?", articleId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *commentDao) GetCommentsByUserId(userId uint) ([]model.Comment, error) {
	db := core.GetDb()
	var comments []model.Comment
	if err := db.Where("user_id = ?", userId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *commentDao) UpdateComment(comment model.Comment) (*model.Comment, error) {
	db := core.GetDb()
	if err := db.Save(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
