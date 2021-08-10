package main

import (
	"awesomeProject/nix_project_test/fileRepository"
	"awesomeProject/nix_project_test/model"

	"fmt"
)

func main() {
	userRepository := fileRepository.NewUserFileRepository()

	//fmt.Println(userRepository)

	u := model.User{
		Name:        "Yarik",
		Email:       "gya@gm.com",
		Password:    "paswordsdf",
		Location:    "home19283012x120830131",
		PhoneNumber: "05099922213",
		Deleted:     false,
	}
	//storedUser, err := userRepository.Create(&u)
	//if err !=nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	nextid, _ := userRepository.Create(&u)
	//fmt.Println(storedUser)
	fmt.Println(nextid)
}
