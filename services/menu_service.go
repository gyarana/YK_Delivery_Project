package services

import (
	"nix_education/model"
	"nix_education/model/repositories"
)

func NewProductService(menuRepository repositories.MenuRepositoryI) *MenuService {
	return &MenuService{
		menuRepository: menuRepository,
	}
}

type MenuServiceI interface {
	GetAllMenuByRestID(idRest int) (*[]model.Product, error)
	GetAllMenu() (*[]model.Product, error)
	GetMenuById(idRest, idMenu int) (*model.Product, error)
}

type MenuService struct {
	menuRepository repositories.MenuRepositoryI
}

func (m MenuService) GetAllMenuByRestID(idRest int) (*[]model.Product, error) {
	menu, err := m.menuRepository.GetAllMenuByRest(idRest)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (m MenuService) GetAllMenu() (*[]model.Product, error) {
	products, err := m.menuRepository.GetAllMenu()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (m MenuService) GetMenuById(idRest, idMenu int) (*model.Product, error) {
	menu, err := m.menuRepository.GetMenuByRestID(idMenu, idRest)
	if err != nil {
		return nil, err
	}
	return menu, nil
}
