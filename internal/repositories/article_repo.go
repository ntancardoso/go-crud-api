package repositories

import (
	"go-crud-api/internal/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{DB: db}
}

func (r *ArticleRepository) GetAll() ([]models.Article, error) {
	var articles []models.Article
	if err := r.DB.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *ArticleRepository) GetByID(id uint) (*models.Article, error) {
	var article models.Article
	if err := r.DB.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) Create(article *models.Article) error {
	return r.DB.Create(article).Error
}

/*
func (r *ArticleRepository) Update(id uint, article *models.Article) error {
	return r.DB.Model(&models.Article{}).Where("id = ?", id).Update(article).Error
}*/

func (r *ArticleRepository) Update(id uint, article *models.Article) error {
	return r.DB.Model(&models.Article{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"title": article.Title,
			"body":  article.Body,
		}).Error
}

func (r *ArticleRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Article{}, id).Error
}
