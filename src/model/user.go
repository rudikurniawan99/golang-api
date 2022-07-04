package model

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Email    string `gorm:"type:varchar(40)"`
		Password string `gorm:"type:varchar(65)"`
	}

	UserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	UserUsecase interface {
		RegisterUsecase(user *User) error
		FindUserByEmailUsercase(user *User) error
		ComparePasswordUsercase(hasPassword, password string) error
	}

	UserRepository interface {
		Create(u *User) error
		Fetch(u *User) error
	}
)
