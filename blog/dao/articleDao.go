package dao

import (
	"blog/core"
	"blog/model"
)

type ArticleDao interface {
	CreateArticle(post model.Post) (*model.Post, error)
	GetArticleById(id uint) (*model.Post, error)
	UpdateArticle(post model.Post) (*model.Post, error)
	DeleteArticle(id uint) error
	GetAllArticles() ([]model.Post, error)
	GetArticlesByUserId(userId uint) ([]model.Post, error)
}

type articleDao struct{}

func NewArticleDao() ArticleDao {
	return &articleDao{}
}

func (articleDao *articleDao) CreateArticle(post model.Post) (*model.Post, error) {
	db := core.GetDb()
	if err := db.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (articleDao *articleDao) GetArticleById(id uint) (*model.Post, error) {
	db := core.GetDb()
	var post model.Post
	if err := db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (articleDao *articleDao) UpdateArticle(post model.Post) (*model.Post, error) {
	db := core.GetDb()
	if err := db.Save(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (articleDao *articleDao) DeleteArticle(id uint) error {
	db := core.GetDb()
	if err := db.Delete(&model.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (articleDao *articleDao) GetAllArticles() ([]model.Post, error) {
	db := core.GetDb()
	var posts []model.Post
	if err := db.Preload("Comments").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (articleDao *articleDao) GetArticlesByUserId(userId uint) ([]model.Post, error) {
	db := core.GetDb()
	var posts []model.Post
	if err := db.Preload("Comments").Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
