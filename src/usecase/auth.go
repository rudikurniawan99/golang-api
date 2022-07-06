package usecase

import (
	"api-2/src/model"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository model.UserRepository
}

func (uc *userUsecase) RegisterUsecase(user *model.User) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashPassword)

	err := uc.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUsecase) FindUserByEmailUsecase(user *model.User) error {
	err := uc.userRepository.Fetch(user)

	if err != nil {
		return err
	}

	return nil

}

func (uc *userUsecase) ComparePasswordUsecase(hashPassword, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func NewUserUsecase(ur model.UserRepository) model.UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}
