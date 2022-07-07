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

func NewBlogUsecase(r model.BlogRepository) model.BlogUsecase {
	return &blogUsecase{
		blogRepository: r,
	}
}
