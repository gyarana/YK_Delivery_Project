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
	GetMenuById(idMenu int) (*model.Product, error)
	CreateMenu(menu *model.ProductParse) error
	UpdateMenu(product *model.ProductParse) error
	DeleteMenu(idMenu int) error
}

type MenuService struct {
	menuRepository repositories.MenuRepositoryI
}

func (m MenuService) UpdateMenu(product *model.ProductParse) error {
	err := m.menuRepository.UpdateMenu(product)
	if err != nil {
		return err
	}
	return nil

}

func (m MenuService) DeleteMenu(idMenu int) error {
	err := m.menuRepository.DeleteMenu(idMenu)
	if err != nil {
		return err
	}
	return nil
}

func (m MenuService) CreateMenu(menu *model.ProductParse) error {
	err := m.menuRepository.CreateMenu(menu)
	if err != nil {
		return err
	}
	return nil
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

func (m MenuService) GetMenuById(idMenu int) (*model.Product, error) {
	menu, err := m.menuRepository.GetMenuByID(idMenu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}
