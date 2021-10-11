package services

import (
	"errors"
	"nix_education/model"
	"nix_education/model/repositories"
)

func NewSupplierService(supplierRepository repositories.RestaurantsRepositoryI) *SupplierService {
	return &SupplierService{
		supplierRepository,
	}
}

type SupplierServiceI interface {
	GetByID(idRest int) (*model.Restaurant, error)
	GetAll() (*[]model.Restaurant, error)
	GetAllByType(restType string) (*[]model.Restaurant, error)
}

type SupplierService struct {
	supplierRepository repositories.RestaurantsRepositoryI
}

func (s SupplierService) GetByID(idRest int) (*model.Restaurant, error) {
	supplier, err := s.supplierRepository.GetSuppliersByID(idRest)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s SupplierService) GetAll() (*[]model.Restaurant, error) {
	suppliers, err := s.supplierRepository.GetAllSuppliers()
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s SupplierService) GetAllByType(restType string) (*[]model.Restaurant, error) {
	suppliers, err := s.supplierRepository.GetSuppliersByType(restType)
	if err != nil {
		return nil, err
	}
	if suppliers == nil {
		return nil, errors.New("no suppliers of this type")
	}
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}
