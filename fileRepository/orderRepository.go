package fileRepository

import "awesomeProject/nix_project_test/model"

type OrderRepositoryI interface {
	Create(u *model.Order) (*model.Order, error)
	Get(email *string, id *int32) *model.Order
	GetAll() []model.Order
	Delete(id int32) error
	Edit(u model.Order) error
}

type OrderDBRepository struct {
}

func (o OrderDBRepository) Create(u *model.Order) (*model.Order, error) {
	panic("implement me")
}

func (o OrderDBRepository) Get(email *string, id *int32) *model.Order {
	panic("implement me")
}

func (o OrderDBRepository) GetAll() []model.Order {
	panic("implement me")
}

func (o OrderDBRepository) Delete(id int32) error {
	panic("implement me")
}

func (o OrderDBRepository) Edit(u model.Order) error {
	panic("implement me")
}
