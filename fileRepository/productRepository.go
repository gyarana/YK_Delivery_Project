package fileRepository

import "awesomeProject/nix_project_test/model"

type ProductRepositoryI interface {
	Create(u *model.Product) (*model.Product, error)
	Get(email *string, id *int32) *model.Product
	GetAll() []model.Product
	Delete(id int32) error
	Edit(u model.Product) error
}

type ProductDBRepository struct {
}

func (p ProductDBRepository) Create(u *model.Product) (*model.Product, error) {
	panic("implement me")
}

func (p ProductDBRepository) Get(email *string, id *int32) *model.Product {
	panic("implement me")
}

func (p ProductDBRepository) GetAll() []model.Product {
	panic("implement me")
}

func (p ProductDBRepository) Delete(id int32) error {
	panic("implement me")
}

func (p ProductDBRepository) Edit(u model.Product) error {
	panic("implement me")
}
