package service

import (
	"blog/dao"
	"blog/model"
)

type CommentService interface {
	CreateComment(comment model.Comment) (*model.Comment, error)

	DeleteComment(id uint, userId uint) error

	GetCommentById(id uint) (*model.Comment, error)

	GetCommentsByArticleId(articleId uint) ([]model.Comment, error)

	GetCommentsByUserId(userId uint) ([]model.Comment, error)

	UpdateComment(comment model.Comment) (*model.Comment, error)
}

type commentService struct {
	commentDao dao.CommentDao
}

func NewCommentService(c dao.CommentDao) CommentService {
	return &commentService{
		commentDao: c,
	}
}

func (c *commentService) CreateComment(comment model.Comment) (*model.Comment, error) {
	if comment, err := c.commentDao.CreateComment(comment); err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}
func (c *commentService) DeleteComment(id uint, userId uint) error {
	if err := c.commentDao.DeleteComment(id, userId); err != nil {
		return err
	}
	return nil
}
func (c *commentService) GetCommentById(id uint) (*model.Comment, error) {
	if comment, err := c.commentDao.GetCommentById(id); err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}
func (c *commentService) GetCommentsByArticleId(articleId uint) ([]model.Comment, error) {
	if comments, err := c.commentDao.GetCommentsByArticleId(articleId); err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}
func (c *commentService) GetCommentsByUserId(userId uint) ([]model.Comment, error) {
	if comments, err := c.commentDao.GetCommentsByUserId(userId); err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}
func (c *commentService) UpdateComment(comment model.Comment) (*model.Comment, error) {
	if comment, err := c.commentDao.UpdateComment(comment); err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}
