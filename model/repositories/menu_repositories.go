package repositories

import (
	"database/sql"
	"nix_education/model"
	"time"
)

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

type MenuRepositoryI interface {
	CreateMenu(idRest int, menu *model.Product) error
	GetMenuByRestID(idMenu, idRest int) (*model.Product, error)
	GetAllMenuByRest(idRest int) (*[]model.RestarauntMenu, error)
	UpdateMenu(idRest int, product *model.Product) error
	DeleteMenu(idMenu, idRest int) error
}

type MenuRepository struct {
	db *sql.DB
}

func (r MenuRepository) CreateMenu(idRest int, product *model.Product) error {
	//	var ing []string = product.Ingredients
	_, err := r.db.Exec("INSERT INTO product ( id,name, image,price, type, ingredients, created_date,id_supplier) VALUES (?,?,?,?,?,?,?,?)", product.ID, product.Name, product.Image, product.Price, product.Type, product.Ingredients, time.Now(), idRest)
	if err != nil {
		return err
	}
	return nil
}

func (r MenuRepository) GetMenuByRestID(idMenu, idRest int) (*model.Product, error) {
	rows, err := r.db.Query("select * from product where id_supplier = ? and id=?", idRest, idMenu)
	if err != nil {
		return nil, err
	}
	var product model.Product

	for rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Image, &product.Price, &product.Type, &product.Ingredients, &product.CreatedDate, &product.UpdatedDate, &product.DeletedDate, &product.IsDeleted, &product.IDSupplier)
	}
	return &product, nil
}

func (r MenuRepository) GetAllMenuByRest(idRest int) (*[]model.RestarauntMenu, error) {
	rows, err := r.db.Query("select * from product where id_supplier=?", idRest)
	if err != nil {
		return nil, err
	}
	var menu []model.RestarauntMenu
	for rows.Next() {
		var menuItems model.RestarauntMenu
		rows.Scan()
		menu = append(menu, menuItems)
	}
	return &menu, nil
}

func (r MenuRepository) UpdateMenu(idRest int, product *model.Product) error {
	_, err := r.db.Exec("UPDATE product SET name = ?, image = ?, type=?, ingredients=?, updated_date=? WHERE id_supplier =? and id = ?", product.Name, product.Image, product.Type, product.Ingredients, time.Now(), idRest, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r MenuRepository) DeleteMenu(idMenu, idRest int) error {
	_, err := r.db.Exec("UPDATE product SET deleted_date=?, is_deleted=? WHERE id_supplier =? and id = ?", time.Now(), true, idRest, idMenu)
	if err != nil {
		return err
	}
	return nil
}
