package fileRepository

import "awesomeProject/nix_project_test/model"

type CartRepositoryI interface {
	Create(u *model.Cart) (*model.Cart, error)
	Get(email *string, id *int32) *model.Cart
	GetAll() []model.Cart
	Delete(id int32) error
	Edit(u model.Cart) error
}

type CartDBRepository struct {
}

func (c CartDBRepository) Create(u *model.Cart) (*model.Cart, error) {
	panic("implement me")
}

func (c CartDBRepository) Get(email *string, id *int32) *model.Cart {
	panic("implement me")
}

func (c CartDBRepository) GetAll() []model.Cart {
	panic("implement me")
}

func (c CartDBRepository) Delete(id int32) error {
	panic("implement me")
}

func (c CartDBRepository) Edit(u model.Cart) error {
	panic("implement me")
}
