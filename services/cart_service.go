package services

import (
	"nix_education/model"
	"nix_education/model/repositories"
)

func NewCartService(cartRepository repositories.CartRepositoryI) *CartService {
	return &CartService{
		cartRepository: cartRepository,
	}
}

type CartServiceI interface {
	CreateCart(cart *model.Cart) error
	GetCartByID(idCart int) (*model.Cart, error)
	UpdateCart(cart *model.Cart) error
	DeleteCart(idCart int) error
}

type CartService struct {
	cartRepository repositories.CartRepositoryI
}

func (c CartService) CreateCart(cart *model.Cart) error {
	err := c.cartRepository.CreateCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func (c CartService) GetCartByID(idCart int) (*model.Cart, error) {
	cart, err := c.cartRepository.GetCartByID(idCart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c CartService) UpdateCart(cart *model.Cart) error {
	err := c.cartRepository.UpdateCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func (c CartService) DeleteCart(idCart int) error {
	err := c.cartRepository.DeleteCart(idCart)
	if err != nil {
		return err
	}
	return nil
}
