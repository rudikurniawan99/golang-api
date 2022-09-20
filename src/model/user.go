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
		Password string `json:"password" validate:"required,min=8"`
	}

	UserUsecase interface {
		RegisterUsecase(user *User) error
		FindUserByEmailUsecase(user *User) error
		FindUserById(user *User, id string) error
		ComparePasswordUsecase(hasPassword, password string) error
		GetAllUserUsecase(user *[]User) error
	}

	UserRepository interface {
		Create(u *User) error
		Fetch(u *User) error
		GetById(u *User, id string) error
		GetAll(u *[]User) error
	}
)
