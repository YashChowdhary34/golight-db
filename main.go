package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string
	State string
	Company string
	Pincode json.Number
}

type User struct {
	Name string
	Age  json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err := nil {
		fmt.Println("Error", err)
	}

	employees := []User {
		{"John", "23", "0123456789", "XYZ Tech", Address{"banglore", "karnataka", "India", "410013"}},
		{"Paul", "21", "0123456789", "Info Tech", Address{"san francisco", "california", "USA", "21444"}},
		{"Robert", "32", "0123456789"}
	}

	for _, value := range employees{
		db.Write("users", value.Name, User{
			Name: value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err := nil{
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records{
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
}