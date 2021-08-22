package DBrepositories

import (
	"database/sql"
	"nix_education/pkg/model"
)

type UserDBRepositoryI interface {
	CreateUser(u *model.User) error
	GetUser(id int32) (*model.User, error)
	GetAllUsers() (*[]model.User, error)
	DeleteUser(id int32) error
	EditUser(u *model.User) error
}

type UserDBRepository struct {
	db *sql.DB
}

func NewUserDBRepository(DB *sql.DB) *UserDBRepository {
	return &UserDBRepository{
		db: DB,
	}
}

func (udbr UserDBRepository) CreateUser(u *model.User) error {

	_, err := udbr.db.Exec("INSERT INTO users ( name, email, password, phonenumber) VALUES (?,?,?,?)", u.Name, u.Email, u.Password, u.PhoneNumber)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (udbr UserDBRepository) GetUser(id int32) (*model.User, error) {
	rows, err := udbr.db.Query("select * from users where id = ?", id)
	if err != nil {
		return &model.User{}, err
	} else {
		var user model.User
		for rows.Next() {
			rows.Scan(&user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.ID)
		}
		return &user, nil
	}
}

func (udbr UserDBRepository) GetAllUsers() (*[]model.User, error) {
	rows, err := udbr.db.Query("select * from users")
	if err != nil {
		return &[]model.User{}, err
	} else {
		var users []model.User
		for rows.Next() {
			var user model.User
			rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.PhoneNumber)
			users = append(users, user)
		}
		return &users, nil
	}
}

func (udbr UserDBRepository) DeleteUser(id int32) error {
	_, err := udbr.db.Exec("delete from users where id=?", id)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (udbr UserDBRepository) EditUser(u *model.User) error {
	_, err := udbr.db.Exec("UPDATE users SET name = ?, email = ?, password = ?, phonenumber = ? WHERE id =?", u.Name, u.Email, u.Password, u.PhoneNumber, u.ID)
	if err != nil {
		return err
	} else {
		return nil
	}
}
