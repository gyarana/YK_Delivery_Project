package fileRepository

import (
	"awesomeProject/nix_project_test/helper"
	"awesomeProject/nix_project_test/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type UserRepositoryI interface {
	Create(u *model.User) (*model.User, error)
	Get(email *string, id *int32) *model.User
	GetAll() []model.User
	Delete(id int32) error
	Edit(u model.User) error
}

type UserDBRepository struct {
}

func (u2 UserDBRepository) Create(u *model.User) (*model.User, error) {
	panic("implement me")
}

func (u2 UserDBRepository) Get(email *string, id *int32) *model.User {
	panic("implement me")
}

func (u2 UserDBRepository) GetAll() []model.User {
	panic("implement me")
}

func (u2 UserDBRepository) Delete(id int32) error {
	panic("implement me")
}

func (u2 UserDBRepository) Edit(u model.User) error {
	panic("implement me")
}

type UserFileRepository struct {
	idMutex *sync.Mutex
}

var (
	users []model.User
)

func NewUserFileRepository() *UserFileRepository {
	return &UserFileRepository{
		idMutex: &sync.Mutex{},
	}
}

func (ufr *UserFileRepository) Create(user *model.User) (*model.User, error) {

	user.ID, _ = ufr.GetNextID()
	err := helper.CreateModel("user", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr *UserFileRepository) Get(email string, id int32) (*model.User, error) {

	users := helper.OpenFilaAndMarshalData()

	for _, user := range users {
		if user.ID == id && user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User is not found. Please check input data")

}

func (ufr *UserFileRepository) GetAll() ([]model.User, error) {

	users := helper.OpenFilaAndMarshalData()

	if len(users) < 1 {
		return nil, fmt.Errorf("users file is empty")
	}

	return users, nil

}

func (ufr *UserFileRepository) Delete(id int32) error {

	users := helper.OpenFilaAndMarshalData()

	for i, user := range users {
		if user.ID == id {
			users[i].Deleted = true
			break
		}
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		panic("")
	}

	err = ioutil.WriteFile("./nix_project_test/datastore/user.txt", bytes, 0600)
	if err != nil {
		panic(" ")
	}

	fmt.Println("User account successfully deleted")
	return nil
}

func (ufr *UserFileRepository) Edit(user *model.User) error {
	var isUpdate bool
	users := helper.OpenFilaAndMarshalData()

	for i, u := range users {
		if u.ID == user.ID {
			users[i] = *user

			bytes, err := json.Marshal(users)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = ioutil.WriteFile("./nix_project_test/datastore/user.txt", bytes, 0600)
			if err != nil {
				fmt.Println(err.Error())
			}
			isUpdate = true
		}
	}
	if isUpdate {
		fmt.Println("Updating is successful")
		return nil
	} else {
		return fmt.Errorf("User is not found. Please check input data")
	}

}

func (ufr *UserFileRepository) GetNextID() (int32, error) {
	ufr.idMutex.Lock()
	users := helper.OpenFilaAndMarshalData()
	ufr.idMutex.Unlock()
	return int32(len(users)), nil

}
