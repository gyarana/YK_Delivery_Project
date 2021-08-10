package fileRepository

import "awesomeProject/nix_project_test/model"

type SuplierRepositoryI interface {
	Create(u *model.Supplier) (*model.Supplier, error)
	Get(email *string, id *int32) *model.Supplier
	GetAll() []model.Supplier
	Delete(id int32) error
	Edit(u model.Supplier) error
}

type SuplierDBRepository struct {
}

func (s SuplierDBRepository) Create(u *model.Supplier) (*model.Supplier, error) {
	panic("implement me")
}

func (s SuplierDBRepository) Get(email *string, id *int32) *model.Supplier {
	panic("implement me")
}

func (s SuplierDBRepository) GetAll() []model.Supplier {
	panic("implement me")
}

func (s SuplierDBRepository) Delete(id int32) error {
	panic("implement me")
}

func (s SuplierDBRepository) Edit(u model.Supplier) error {
	panic("implement me")
}
