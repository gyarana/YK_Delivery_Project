package services

import (
	"nix_education/model"
	"nix_education/model/repositories"
)

func NewOrderService(orderRepository repositories.OrderRepositoryI) *OrderService {
	return &OrderService{
		orderRepository,
	}
}

type OrderServiceI interface {
	CreateOrder(u *model.OrderRequest) error
	GetOrder(id int32) (*model.Order, error)
	GetAllOrders() (*[]model.Order, error)
	DeleteOrder(id int32) error
	EditOrder(u *model.OrderRequest) error
}

type OrderService struct {
	orderRepository repositories.OrderRepositoryI
}

func (o OrderService) CreateOrder(u *model.OrderRequest) error {
	err := o.orderRepository.CreateOrder(u)
	if err != nil {
		return err
	}
	return nil
}

func (o OrderService) GetOrder(id int32) (*model.Order, error) {
	order, err := o.orderRepository.GetOrder(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o OrderService) GetAllOrders() (*[]model.Order, error) {
	orders, err := o.orderRepository.GetAllOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o OrderService) DeleteOrder(id int32) error {
	err := o.orderRepository.DeleteOrder(id)
	if err != nil {
		return err
	}
	return nil
}

func (o OrderService) EditOrder(u *model.OrderRequest) error {
	err := o.orderRepository.EditOrder(u)
	if err != nil {
		return err
	}
	return nil
}
