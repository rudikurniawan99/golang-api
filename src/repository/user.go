package repository

import (
	"api-2/src/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	err := r.db.Create(user).Error

	return err
}

func (r *userRepository) Fetch(user *model.User) error {
	err := r.db.Where(&model.User{Email: user.Email}).First(user).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetById(user *model.User, id string) error {
	err := r.db.First(&user, id).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAll(user *[]model.User) error {
	err := r.db.Find(&user).Error

	if err != nil {
		return err
	}
	return nil
}

// func (r *userRepository) GetByEmail(email string) (model.User, error) {
// 	user := model.User{}

// 	err := r.db.Where(&model.User{Email: email}).First(&user).Error

// 	if err != nil {
// 		return user, err
// 	} else {
// 		fmt.Println(user)
// 		return user, nil
// 	}
// }
