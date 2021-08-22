package DBrepositories

import (
	"database/sql"
	"nix_education/pkg/model"
)

type OrderDBRepositoryI interface {
	CreateOrder(u *model.Order) error
	GetOrder(id int32) (*model.Order, error)
	GetAllOrders() (*[]model.Order, error)
	DeleteOrder(id int32) error
	EditOrder(u *model.Order) error
}

type OrderDBRepository struct {
	db *sql.DB
}

func NewOrderDBRepository(DB *sql.DB) *OrderDBRepository {
	return &OrderDBRepository{
		db: DB,
	}
}

func (odbr OrderDBRepository) CreateOrder(o *model.Order) error {

	_, err := odbr.db.Exec("INSERT INTO orders ( entities, status, adress) VALUES (?,?,?)", o.Entities, o.Status, o.Adress)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (odbr OrderDBRepository) GetOrder(id int32) (*model.Order, error) {
	rows, err := odbr.db.Query("select * from orders where id = ?", id)
	if err != nil {
		return &model.Order{}, err
	} else {
		var order model.Order
		for rows.Next() {
			rows.Scan(&order.Id, &order.Entities, &order.Status, &order.Adress)
		}
		return &order, nil
	}
}

func (odbr OrderDBRepository) GetAllOrders() (*[]model.Order, error) {
	rows, err := odbr.db.Query("select * from orders")
	if err != nil {
		return &[]model.Order{}, err
	} else {
		var orders []model.Order
		for rows.Next() {
			var order model.Order
			rows.Scan(&order.Id, &order.Entities, &order.Status, &order.Adress)
			orders = append(orders, order)
		}
		return &orders, nil
	}
}

func (odbr OrderDBRepository) DeleteOrder(id int32) error {
	_, err := odbr.db.Exec("delete from orders where id=?", id)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (odbr OrderDBRepository) EditOrder(o *model.Order) error {
	_, err := odbr.db.Exec("UPDATE orders SET entities = ?, status = ?, adress = ? WHERE id =?", o.Entities, o.Status, o.Adress)
	if err != nil {
		return err
	} else {
		return nil
	}
}
