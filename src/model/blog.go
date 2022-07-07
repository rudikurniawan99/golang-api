package model

import "gorm.io/gorm"

type (
	Blog struct {
		gorm.Model
		Title  string `gorm:"type:varchar(250)"`
		Body   string `gorm:"type:text"`
		UserID uint
		User   User
	}
	BlogRequest struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	BlogUsecase interface {
		CreateUsecase(blog *Blog) error
		DeleteUsecase(id int) error
		FindByIdUsecase(blog *Blog, id int) error
	}

	BlogRepository interface {
		Create(blog *Blog) error
		Delete(id int) error
		FindById(blog *Blog, id int) error
	}
)
