package main

import (
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/repository/mysql"
)

func main() {

	testUserRepo();
}

func testUserRepo() {
	repo := mysql.New()
	created_user, err := repo.Register(entity.User{
		ID:          0,
		Name:        "abbas",
		PhoneNumber: "09356743672",
	})
	if err != nil {
		fmt.Println("register error :", err)
	} else {
		fmt.Println("user created successfully", created_user)
	}

	unique, err := repo.IsPhoneNumberUnique(created_user.PhoneNumber)
	if err != nil {
		fmt.Println("phone number is not unique", err)
	} else {
		fmt.Println("phone number is untie", unique)
	}
}
