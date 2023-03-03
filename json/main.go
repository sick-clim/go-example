package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string     `json:"name"`
	Age     int        `json:"age"`
	Email   string     `json:"email"`
	Address []*Address `json:"address,omitempty"`
}

type PersonNotPointer struct {
	Name    string     `json:"name"`
	Age     int        `json:"age"`
	Email   string     `json:"email"`
	Address []Address `json:"address,omitempty"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	person := &Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Address: []*Address{
			{
				City:    "Tokyo",
				Country: "Japan",
			},
			{
				City:    "New York",
				Country: "USA",
			},
		},
	}

	data, _ := json.Marshal(person)
	fmt.Println(string(data)) // {"name":"Alice","age":30,"email":"alice@example.com",
	// "address":[{"city":"Tokyo","country":"Japan"},{"city":"New York","country":"USA"}]}
	person2 := &Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Printf("%T\n", person2.Address)
	data2, _ := json.Marshal(person2)
	fmt.Println(string(data2)) // {"name":"Alice","age":30,"email":"alice@example.com",

	personNotP := PersonNotPointer{
		Name:  "Jhon",
		Age:   42,
		Email: "jhon@example.com",
	}
	fmt.Printf("%T\n", personNotP.Address)
	data3, _ := json.Marshal(personNotP)
	fmt.Println(string(data3))


}
