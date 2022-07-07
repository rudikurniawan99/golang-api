package usecase

import (
	"api-2/src/model"
)

type blogUsecase struct {
	blogRepository model.BlogRepository
}

func (u *blogUsecase) CreateUsecase(blog *model.Blog) error {
	err := u.blogRepository.Create(blog)

	if err != nil {
		return err
	}

	return nil
}

func (u *blogUsecase) FindByIdUsecase(blog *model.Blog, id int) error {
	err := u.blogRepository.FindById(blog, id)

	if err != nil {
		return err
	}
	return nil
}

func (u *blogUsecase) DeleteUsecase(id int) error {
	err := u.blogRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func NewBlogUsecase(r model.BlogRepository) model.BlogUsecase {
	return &blogUsecase{
		blogRepository: r,
	}
}
