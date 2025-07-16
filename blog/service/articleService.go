package service

import (
	"blog/dao"
	"blog/model"
)

type ArticleService interface {
	CreateArticle(post model.Post) (*model.Post, error)
	GetArticleById(id uint) (*model.Post, error)
	UpdateArticle(post model.Post) (*model.Post, error)
	DeleteArticle(id uint) error
	GetArticlesByUserId(userId uint) ([]model.Post, error)
	GetAllArticles() ([]model.Post, error)
}

type articleService struct {
	articleDao dao.ArticleDao
}

func NewArticleService(dao dao.ArticleDao) ArticleService {
	return &articleService{
		articleDao: dao,
	}
}

func (articleService *articleService) CreateArticle(post model.Post) (*model.Post, error) {
	if post, err := articleService.articleDao.CreateArticle(post); err != nil {
		return nil, err
	} else {
		return post, nil
	}
}

func (articleService *articleService) GetArticleById(id uint) (*model.Post, error) {
	if post, err := articleService.articleDao.GetArticleById(id); err != nil {
		return nil, err
	} else {
		return post, nil
	}
}

func (articleService *articleService) UpdateArticle(post model.Post) (*model.Post, error) {
	if post, err := articleService.articleDao.UpdateArticle(post); err != nil {
		return nil, err
	} else {
		return post, nil
	}
}

func (articleService *articleService) DeleteArticle(id uint) error {
	if err := articleService.articleDao.DeleteArticle(id); err != nil {
		return err
	}
	return nil
}

func (articleService *articleService) GetArticlesByUserId(userId uint) ([]model.Post, error) {
	if posts, err := articleService.articleDao.GetArticlesByUserId(userId); err != nil {
		return nil, err
	} else {
		return posts, nil
	}
}

func (articleService *articleService) GetAllArticles() ([]model.Post, error) {
	if posts, err := articleService.articleDao.GetAllArticles(); err != nil {
		return nil, err
	} else {
		return posts, nil
	}
}
