package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	username string
	age      int
}

func (u *User) Create(name string, age int) error {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()

	if err := db.Create(u{name: "ahmed", "age": 12}).Error; err != nil {
		return err
	} else {
		return nil
	}

}

func (*User) GetAll() ([]User, error) {
	db, _ := gorm.Open("sqlite3", "test.db")
	var user []User
	err := db.Order("id desc").Limit(3).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	var user User
	db.AutoMigrate(&user)
	err = user.Create("ahmed3", 2320)
	if err != nil {
		fmt.Printf("There is an error: %v", err)
		return
	}
	listUsers, err := user.GetAll()
	if err != nil {
		fmt.Printf("There is an error in getting Db: %v", err)
		return
	}
	for k, v := range listUsers {
		fmt.Printf("The key is: %v, the value is: %v\n", k, v)
	}

}
