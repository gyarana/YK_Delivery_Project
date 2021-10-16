package repositories

import (
	"database/sql"
	"nix_education/model"
	"time"
)

type UserRepositoryI interface {
	CreateUser(u *model.User) error
	GetUser(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
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

	_, err := udbr.db.Exec("INSERT INTO users ( name, email, password, phonenumber,created_date) VALUES (?,?,?,?,?)", u.Name, u.Email, u.PasswordHash, u.PhoneNumber, u.CreatedDate)
	if err != nil {
		return err
	}
	return nil
}

func (udbr UserRepository) GetUser(id int) (*model.User, error) {
	rows, err := udbr.db.Query("select * from users where id = ?", id)
	if err != nil {
		return &model.User{}, err
	}
	var user model.User
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.PhoneNumber, &user.CreatedDate, &user.UpdatedDate, &user.DeletedDate, &user.IsDeleted)
	}
	return &user, nil
}

func (udbr UserRepository) GetUserByEmail(email string) (*model.User, error) {
	rows, err := udbr.db.Query("select * from users where email = ?", email)
	if err != nil {
		return &model.User{}, err
	}
	var user model.User
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.PhoneNumber, &user.CreatedDate, &user.UpdatedDate, &user.DeletedDate, &user.IsDeleted)
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
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.PhoneNumber, &user.CreatedDate, &user.UpdatedDate, &user.DeletedDate, &user.IsDeleted)
		users = append(users, user)
	}
	return &users, nil
}

func (udbr UserRepository) DeleteUser(id int32) error {
	_, err := udbr.db.Exec("UPDATE users SET deleted_date=?, is_deleted=? where id=?", time.Now(), true, id)
	if err != nil {
		return err
	}
	return nil
}

func (udbr UserRepository) EditUser(u *model.User) error {
	_, err := udbr.db.Exec("UPDATE users SET name = ?, email = ?, password_hash = ?, phone_number = ?, created_date = ?, updated_date = ?, deleted_date = ?, is_deleted = ? WHERE id =?", u.Name, u.Email, u.PasswordHash, u.PhoneNumber, u.CreatedDate, u.UpdatedDate, u.DeletedDate, u.IsDeleted, u.ID)
	if err != nil {
		return err
	}
	return nil
}
