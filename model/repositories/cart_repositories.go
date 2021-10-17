package repositories

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"nix_education/model"
	"time"
)

func NewCartRepository(db *sql.DB, logger *logrus.Logger) *CartRepository {
	return &CartRepository{
		db:     db,
		logger: logger,
	}
}

type CartRepositoryI interface {
	CreateCart(cart *model.Cart) error
	GetCartByID(idCart int) (*model.Cart, error)
	UpdateCart(cart *model.Cart) error
	DeleteCart(idCart int) error
}

type CartRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func (c CartRepository) CreateCart(cart *model.Cart) error {
	_, err := c.db.Exec("INSERT INTO cart2 ( created_date) VALUES (?)", time.Now().Format(time.RFC1123))
	if err != nil {
		return err
	}

	for _, product := range cart.Products {
		_, err := c.db.Exec("INSERT INTO cart_products ( cartID, productID, quantity) VALUES (?,?,?)", cart.ID, product.ProductID, product.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c CartRepository) GetCartByID(idCart int) (*model.Cart, error) {
	rows, err := c.db.Query("select * from cart2 where id=?", idCart)
	if err != nil {
		return nil, err
	}
	var cart model.Cart

	for rows.Next() {
		rows.Scan(&cart.ID, &cart.Products, &cart.CreatedDate, &cart.UpdatedDate, &cart.DeletedDate, &cart.IsDeleted)
	}
	return &cart, nil
}

func (c CartRepository) UpdateCart(cart *model.Cart) error {

	for _, product := range cart.Products {
		_, err := c.db.Exec("UPDATE  cart_products SET productID = ?, quantity = ? WHERE cartID=?", product.ProductID, product.Quantity, cart.ID)
		if err != nil {
			return err
		}
	}
	_, err := c.db.Exec("UPDATE  cart2 SET updated_date = ?  WHERE id=? ", time.Now().Format(time.RFC1123), cart.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c CartRepository) DeleteCart(idCart int) error {
	_, err := c.db.Exec("UPDATE cart SET deleted_date=?, is_deleted=? WHERE id = ?", time.Now().Format(time.RFC1123), true, idCart)
	if err != nil {
		return err
	}
	return nil
}
