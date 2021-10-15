package repositories

import (
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"nix_education/model"
	"time"
)

func NewMenuRepository(db *sql.DB, logger *logrus.Logger) *MenuRepository {
	return &MenuRepository{
		db:     db,
		logger: logger,
	}
}

type MenuRepositoryI interface {
	CreateMenu(menu *model.ProductParse) error
	GetMenuByID(idMenu int) (*model.Product, error)
	GetAllMenuByRest(idRest int) (*[]model.Product, error)
	UpdateMenu(product *model.ProductParse) error
	DeleteMenu(idMenu int) error
	GetAllMenu() (*[]model.Product, error)
}

type MenuRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func (r MenuRepository) CreateMenu(product *model.ProductParse) error {
	ing, err := json.MarshalIndent(product.Ingredients, "", "")
	if err != nil {
		r.logger.Error("We have some problem with unmarshalling ingredients.Please check it!")
	}
	_, err = r.db.Exec("INSERT INTO product ( id,name, image,price, type, ingredients, created_date,id_supplier) VALUES (?,?,?,?,?,?,?,?)", product.ID, product.Name, product.Image, product.Price, product.Type, ing, time.Now(), product.IDSupplier)
	if err != nil {
		return err
	}
	return nil
}

func (r MenuRepository) GetMenuByID(idMenu int) (*model.Product, error) {
	rows, err := r.db.Query("select * from product where id=?", idMenu)
	if err != nil {
		return nil, err
	}
	var product model.Product

	for rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Image, &product.Price, &product.Type, &product.Ingredients, &product.CreatedDate, &product.UpdatedDate, &product.DeletedDate, &product.IsDeleted, &product.IDSupplier)
	}
	return &product, nil
}

func (r MenuRepository) GetAllMenuByRest(idRest int) (*[]model.Product, error) {
	rows, err := r.db.Query("select * from product where id_supplier=?", idRest)
	if err != nil {
		return nil, err
	}
	var menu []model.Product
	for rows.Next() {
		var menuItems model.Product
		rows.Scan(&menuItems.ID, &menuItems.Name, &menuItems.Image, &menuItems.Price, &menuItems.Type, &menuItems.Ingredients, &menuItems.CreatedDate, &menuItems.UpdatedDate, &menuItems.DeletedDate, &menuItems.IsDeleted, &menuItems.IDSupplier)
		menu = append(menu, menuItems)
	}
	return &menu, nil
}

func (r MenuRepository) UpdateMenu(product *model.ProductParse) error {
	ing, err := json.MarshalIndent(product.Ingredients, "", "")
	if err != nil {
		r.logger.Error("We have some problem with unmarshalling ingredients.Please check it!")
	}
	_, err = r.db.Exec("UPDATE product SET name = ?, image = ?,price = ?, type=?, ingredients=?, updated_date=? WHERE id_supplier =? and id = ?", product.Name, product.Image, product.Price, product.Type, ing, time.Now(), product.IDSupplier, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r MenuRepository) DeleteMenu(idMenu int) error {
	_, err := r.db.Exec("UPDATE product SET deleted_date=?, is_deleted=? WHERE id = ?", time.Now(), true, idMenu)
	if err != nil {
		return err
	}
	return nil
}

func (r MenuRepository) GetAllMenu() (*[]model.Product, error) {
	rows, err := r.db.Query("select * from product ")
	if err != nil {
		return nil, err
	}
	var menu []model.Product
	for rows.Next() {
		var menuItems model.Product
		rows.Scan()
		menu = append(menu, menuItems)
	}
	return &menu, nil
}
