package services

import (
	"errors"
	"go-crud-api/internal/models"
	"go-crud-api/internal/repositories"
)

type ArticleService struct {
	repo *repositories.ArticleRepository
}

func NewArticleService(repo *repositories.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) GetAllArticles() ([]models.Article, error) {
	return s.repo.GetAll()
}

func (s *ArticleService) GetArticleByID(id uint) (*models.Article, error) {
	article, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errors.New("article not found")
	}
	return article, nil
}

func (s *ArticleService) CreateArticle(article *models.Article) error {
	return s.repo.Create(article)
}

func (s *ArticleService) UpdateArticle(id uint, article *models.Article) error {
	existingArticle, err := s.GetArticleByID(id)
	if err != nil {
		return err
	}
	if existingArticle == nil {
		return errors.New("article not found")
	}
	return s.repo.Update(id, article)
}

func (s *ArticleService) DeleteArticle(id uint) error {
	existingArticle, err := s.GetArticleByID(id)
	if err != nil {
		return err
	}
	if existingArticle == nil {
		return errors.New("article not found")
	}
	return s.repo.Delete(id)
}
