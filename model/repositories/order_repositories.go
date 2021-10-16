package repositories

import (
	"database/sql"
	"nix_education/model"
	"time"
)

type OrderRepositoryI interface {
	CreateOrder(u *model.OrderRequest) error
	GetOrder(id int32) (*model.Order, error)
	GetAllOrders() (*[]model.Order, error)
	DeleteOrder(id int32) error
	EditOrder(u *model.OrderRequest) error
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(DB *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: DB,
	}
}

func (odbr OrderRepository) CreateOrder(o *model.OrderRequest) error {

	_, err := odbr.db.Exec("INSERT INTO orders ( user_id,cart_id, status,created_date) VALUES (?,?,?,?)", o.UserID, o.CartID, o.Status, time.Now())
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
		rows.Scan(&order.OrderID, &order.UserID, &order.CartID, &order.Status, &order.CreatedDate, &order.UpdatedDate, &order.DeletedDate, &order.IsDeleted)
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
		rows.Scan(&order.OrderID, &order.UserID, &order.CartID, &order.Status, &order.CreatedDate, &order.UpdatedDate, &order.DeletedDate, &order.IsDeleted)
		orders = append(orders, order)
	}
	return &orders, nil
}

func (odbr OrderRepository) DeleteOrder(id int32) error {
	_, err := odbr.db.Exec("UPDATE orders SET deleted_date=?, is_deleted=? where id=?", time.Now(), true, id)
	if err != nil {
		return err
	}
	return nil

}

func (odbr OrderRepository) EditOrder(o *model.OrderRequest) error {
	_, err := odbr.db.Exec("UPDATE orders SET status = ?, updated_date=? WHERE id =?", o.Status, time.Now(), o.OrderID)
	if err != nil {
		return err
	}
	return nil

}
