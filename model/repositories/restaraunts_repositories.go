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
	CreateSuppliers(restaurant *model.RestaurantParse) error
	GetSuppliersByID(id int) (*model.Restaurant, error)
	GetAllSuppliers() (*[]model.Restaurant, error)
	UpdateSuppliers(restaurant *model.RestaurantParse) error
	DeleteSuppliers(id int) error
	GetSuppliersByType(restType string) (*[]model.Restaurant, error)
}

type RestaurantsRepository struct {
	db *sql.DB
}

func (r RestaurantsRepository) CreateSuppliers(restaurant *model.RestaurantParse) error {
	_, err := r.db.Exec("INSERT INTO suppliers ( id, image, name, type,workingHours, created_date) VALUES (?,?,?,?,?,?)", restaurant.Id, restaurant.Image, restaurant.Name, restaurant.Type, restaurant.WorkingHours, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r RestaurantsRepository) GetSuppliersByID(id int) (*model.Restaurant, error) {
	rows, err := r.db.Query("select * from suppliers where id = ?", id)
	if err != nil {
		return nil, err
	}
	var rest model.Restaurant
	for rows.Next() {
		rows.Scan(&rest.Id, &rest.Image, &rest.Name, &rest.Type, &rest.CreatedDate, &rest.UpdatedDate, &rest.DeletedDate, &rest.IsDeleted)
	}
	return &rest, nil
}

func (r RestaurantsRepository) GetAllSuppliers() (*[]model.Restaurant, error) {
	rows, err := r.db.Query("select * from suppliers")
	if err != nil {
		return &[]model.Restaurant{}, err
	}
	var rests []model.Restaurant
	for rows.Next() {
		var rest model.Restaurant
		rows.Scan()
		rests = append(rests, rest)
	}
	return &rests, nil
}

func (r RestaurantsRepository) UpdateSuppliers(restaurant *model.RestaurantParse) error {
	_, err := r.db.Exec("UPDATE suppliers SET name = ?, image = ?,type = ?, workingHours = ?, updated_date=? WHERE id =?", restaurant.Name, restaurant.Image, restaurant.Type, restaurant.WorkingHours, time.Now(), restaurant.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r RestaurantsRepository) DeleteSuppliers(id int) error {
	_, err := r.db.Exec("UPDATE suppliers SET deleted_date=?, is_deleted=? WHERE id =?", time.Now(), true, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RestaurantsRepository) GetSuppliersByType(restType string) (*[]model.Restaurant, error) {
	rows, err := r.db.Query("select * from suppliers where type = ?", restType)
	if err != nil {
		return &[]model.Restaurant{}, err
	}
	var rests []model.Restaurant
	for rows.Next() {
		var rest model.Restaurant
		rows.Scan()
		rests = append(rests, rest)
	}
	return &rests, nil
}
