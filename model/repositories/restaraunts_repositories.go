package repositories

import (
	"database/sql"
	"nix_education/model"
	"time"
)

func NewRestaurantsRepository(db *sql.DB) *RestaurantsRepository {
	return &RestaurantsRepository{db: db}
}

type RestaurantsRepositoryI interface {
	CreateSuppliers(restaurant *model.Restaurant) error
	GetSuppliersByID(id int) (*model.Restaurant, error)
	GetAllSuppliers() (*[]model.Supliers, error)
	UpdateSuppliers(restaurant *model.Restaurant) error
	DeleteSuppliers(id int) error
}

type RestaurantsRepository struct {
	db *sql.DB
}

func (r RestaurantsRepository) CreateSuppliers(restaurant *model.Restaurant) error {
	_, err := r.db.Exec("INSERT INTO suppliers9 ( name, image, created_date) VALUES (?,?,?)", restaurant.Name, restaurant.Image, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r RestaurantsRepository) GetSuppliersByID(id int) (*model.Restaurant, error) {
	rows, err := r.db.Query("select * from suppliers9 where id = ?", id)
	if err != nil {
		return nil, err
	}
	var rest model.Restaurant
	for rows.Next() {
		rows.Scan(&rest.Id, &rest.Name, &rest.Image, &rest.CreatedDate, &rest.UpdatedDate, &rest.DeletedDate, &rest.IsDeleted)
	}
	return &rest, nil
}

func (r RestaurantsRepository) GetAllSuppliers() (*[]model.Supliers, error) {
	rows, err := r.db.Query("select * from suppliers9")
	if err != nil {
		return &[]model.Supliers{}, err
	}
	var rests []model.Supliers
	for rows.Next() {
		var rest model.Supliers
		rows.Scan()
		rests = append(rests, rest)
	}
	return &rests, nil
}

func (r RestaurantsRepository) UpdateSuppliers(restaurant *model.Restaurant) error {
	_, err := r.db.Exec("UPDATE suppliers9 SET name = ?, image = ?, updated_date=? WHERE id =?", restaurant.Name, restaurant.Image, time.Now(), restaurant.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r RestaurantsRepository) DeleteSuppliers(id int) error {
	_, err := r.db.Exec("UPDATE suppliers9 SET deleted_date=?, is_deleted=? WHERE id =?", time.Now(), true, id)
	if err != nil {
		return err
	}
	return nil
}
