package helper

import (
	"awesomeProject/nix_project_test/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type UserIDFile struct {
	ID int32 `json:"id"`
}

var (
	users []model.User
	uID   []UserIDFile
)

func CreateModel(modelName string, v interface{}) error {
	bytes, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fmt.Sprintf("./nix_project_test/datastore/%s.txt", modelName), os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	fStat, err := file.Stat()
	if err != nil {
		return err
	}
	fSize := fStat.Size()

	if fSize < 1 {
		_, err = file.WriteAt([]byte(`[`), 0)
		if err != nil {
			return err
		}
		bytes = append(bytes, ']')
		_, err = file.WriteAt(bytes, fSize+1)
		if err != nil {
			return err
		}
	}
	if fSize > 1 {
		bytes = append(bytes, ']')
		_, err = file.WriteAt([]byte(`, \n`), fSize-1)
		_, err = file.WriteAt(bytes, fSize+1)
		if err != nil {
			return err
		}
	}

	return nil
}

func OpenFilaAndMarshalData() []model.User {
	fileData, err := os.Open("./nix_project_test/datastore/user.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer fileData.Close()
	bFileData, err := ioutil.ReadAll(fileData)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(bFileData, &users)

	if err != nil {
		fmt.Println(err.Error())
	}

	return users

}
