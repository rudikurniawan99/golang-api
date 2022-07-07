package repository

import (
	"api-2/src/model"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func (r *blogRepository) Create(blog *model.Blog) error {
	err := r.db.Create(blog).Error

	if err != nil {
		return err
	}

	return nil
}

func NewBlogRepository(db *gorm.DB) model.BlogRepository {
	return &blogRepository{db}
}
