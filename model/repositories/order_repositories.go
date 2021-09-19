package repositories

import (
	"database/sql"
	"nix_education/model"
)

type OrderRepositoryI interface {
	CreateOrder(u *model.Order) error
	GetOrder(id int32) (*model.Order, error)
	GetAllOrders() (*[]model.Order, error)
	DeleteOrder(id int32) error
	EditOrder(u *model.Order) error
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(DB *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: DB,
	}
}

func (odbr OrderRepository) CreateOrder(o *model.Order) error {

	_, err := odbr.db.Exec("INSERT INTO orders ( entities, status, adress) VALUES (?,?,?)", o.Entities, o.Status, o.Adress)
	if err != nil {
		return err
	}
	return nil

}

func (odbr OrderRepository) GetOrder(id int32) (*model.Order, error) {
	rows, err := odbr.db.Query("select * from orders where id = ?", id)
	if err != nil {
		return &model.Order{}, err
	}
	var order model.Order
	for rows.Next() {
		rows.Scan(&order.Id, &order.Entities, &order.Status, &order.Adress)
	}
	return &order, nil
}

func (odbr OrderRepository) GetAllOrders() (*[]model.Order, error) {
	rows, err := odbr.db.Query("select * from orders")
	if err != nil {
		return &[]model.Order{}, err
	}
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		rows.Scan(&order.Id, &order.Entities, &order.Status, &order.Adress)
		orders = append(orders, order)
	}
	return &orders, nil
}

func (odbr OrderRepository) DeleteOrder(id int32) error {
	_, err := odbr.db.Exec("delete from orders where id=?", id)
	if err != nil {
		return err
	}
	return nil

}

func (odbr OrderRepository) EditOrder(o *model.Order) error {
	_, err := odbr.db.Exec("UPDATE orders SET entities = ?, status = ?, adress = ? WHERE id =?", o.Entities, o.Status, o.Adress)
	if err != nil {
		return err
	}
	return nil

}
