package repositories

import (
	"database/sql"
	"nix_education/model"
)

type UserRepositoryI interface {
	CreateUser(u *model.User) error
	GetUser(id int32) (*model.User, error)
	GetAllUsers() (*[]model.User, error)
	DeleteUser(id int32) error
	EditUser(u *model.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{
		db: DB,
	}
}

func (udbr UserRepository) CreateUser(u *model.User) error {

	_, err := udbr.db.Exec("INSERT INTO users ( name, email, password, phonenumber) VALUES (?,?,?,?)", u.Name, u.Email, u.Password, u.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (udbr UserRepository) GetUser(id int32) (*model.User, error) {
	rows, err := udbr.db.Query("select * from users where id = ?", id)
	if err != nil {
		return &model.User{}, err
	}
	var user model.User
		for rows.Next() {
			rows.Scan(&user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.ID)
		}
		return &user, nil
}

func (udbr UserRepository) GetAllUsers() (*[]model.User, error) {
	rows, err := udbr.db.Query("select * from users")
	if err != nil {
		return &[]model.User{}, err
	}
	var users []model.User
		for rows.Next() {
			var user model.User
			rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.PhoneNumber)
			users = append(users, user)
		}
		return &users, nil
}

func (udbr UserRepository) DeleteUser(id int32) error {
	_, err := udbr.db.Exec("delete from users where id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (udbr UserRepository) EditUser(u *model.User) error {
	_, err := udbr.db.Exec("UPDATE users SET name = ?, email = ?, password = ?, phonenumber = ? WHERE id =?", u.Name, u.Email, u.Password, u.PhoneNumber, u.ID)
	if err != nil {
		return err
	}
	return nil
}
