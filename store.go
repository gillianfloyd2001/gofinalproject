package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

// connects with the database
func init() {
	var err error
	db, err = sql.Open("postgres", "user=home-town-pizza-api dbname=home-town-pizza password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Employees gets all the employees out.
func Employees() (employees []Employee) {
	rows, err := db.Query("SELECT id, name, position from employees")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		employee := Employee{}
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Position)
		if err != nil {
			panic(err)
		}
		employees = append(employees, employee)
	}
	rows.Close()
	return
}

// GetEmployee will get an employee out by id.
func GetEmployee(id int) (employee Employee, err error) {
	employee = Employee{}
	err = db.QueryRow("SELECT * from Employees where id = $1", id).Scan(&employee.Name, &employee.Position)
	return
}

// Create saves the employee described by this struct to the database.
// If the employee has a valid Id (not 0), no work is done because we
// assume the valid Id was provided by the database on some past save
func (employee *Employee) Create() (err error) {
	// return early if the employee already has an Id other than 0
	// In this case, we are assuming the non-0 id means it has been
	// created because PostgreSQL SERIAL starts at 1
	if employee.Id != 0 {
		return
	}

	createEmployeeStatement, err := db.Prepare("INSERT INTO employees (name, position) VALUES ($1, $2) RETURNING id")
	if err != nil {
		panic(err)
	}
	defer createEmployeeStatement.Close()
	err = createEmployeeStatement.QueryRow(employee.Name, employee.Position).Scan(&employee.Id)
	if err != nil {
		panic(err)
	}
	return
}

// Update will update the list of employee and saves this struct to the database.
func (employee *Employee) Update() (err error) {
	_, err = db.Exec("update employees set name = $2, position = $3 where id = $1", employee.Id, employee.Name, employee.Position)
	return
}

// Delete will delete an employee based of the id that is given.
func (employee *Employee) Delete() (err error) {
	_, err = db.Exec("delete from employees where id = $1", employee.Id)
	return
}

// Deliveries get all the deliveries.
func Deliveries() (deliveries []Delivery) {
	rows, err := db.Query("SELECT id, address, tip from deliveries")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		delivery := Delivery{}
		err = rows.Scan(&delivery.Id, &delivery.Address, &delivery.Tip)
		if err != nil {
			panic(err)
		}
		deliveries = append(deliveries, delivery)
	}
	rows.Close()
	return
}

// GetDelivery gets delivery by an id.
func GetDelivery(id int) (delivery Delivery, err error) {
	delivery = Delivery{}
	err = db.QueryRow("SELECT id, address, tip FROM deliveries WHERE id = $1", id).Scan(&delivery.Id, &delivery.Address, &delivery.Tip)
	return
}

// Create saves the delivery describes this struct to the database.
// If the delivery does not have a vaild Address, no work is done because we
// assume the valid Address was provided by the database on some past save.
func (delivery *Delivery) Create() (err error) {
	// returns early if the address is empty
	if len(delivery.Address) == 0 {
		panic("Delivery created without an address")
	}
	createDeliveryStatement, err := db.Prepare("INSERT INTO deliveries (address, tip) VALUES ($1, $2) RETURNING id")
	if err != nil {
		panic(err)
	}
	defer createDeliveryStatement.Close()
	err = createDeliveryStatement.QueryRow(&delivery.Address, &delivery.Tip).Scan(&delivery.Id)
	if err != nil {
		panic(err)
	}
	return
}
