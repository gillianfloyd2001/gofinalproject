package main

// This creates an Employee
type Employee struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

// This creates a Delivery
type Delivery struct {
	Id          uint   `json:"id"`
	Name        string `json: "name"`
	PhoneNumber string `json: "phoneNumber"`
	Address     string `json:"address"`
	Tip         bool   `json:"tip"`
}
