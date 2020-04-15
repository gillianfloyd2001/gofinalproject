package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func listEmployees(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, Employees())
}

func singleEmployee(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	employee, err := GetEmployee(id)
	return c.JSON(http.StatusOK, employee)
}

func createEmployee(c echo.Context) (err error) {
	employee := new(Employee)
	if err = c.Bind(employee); err != nil {
		return
	}
	createdEmployee := employee.CreateEmployee()
	return c.JSON(http.StatusCreated, createdEmployee)
}

func listDeliveries(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, Deliveries())
}

func singleDelivery(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	delivery, err := GetDelivery(id)
	return c.JSON(http.StatusOK, delivery)
}

func createDelivery(c echo.Context) (err error) {
	delivery := Delivery{}
	if err = c.Bind(delivery); err != nil {
		return
	}
	createdDelivery := delivery.CreateDelivery()
	return c.JSON(http.StatusCreated, createdDelivery)
}

func main() {
	// http.HandleFunc("/clockin", createClockIn)
	// http.HandleFunc("/employees", createListEmployees)
	// http.HandleFunc("/clockout", createClockOut)

	e := echo.New()
	e.GET("/employees", listEmployees)
	e.GET("/employee/:id", singleEmployee)
	e.POST("/employee", createEmployee)

	e.GET("/deliveries", listDeliveries)
	e.GET("/delivery/:id", singleDelivery)
	e.POST("/deliveries", createDelivery)

	fmt.Println("Listening on 8080")
	if err := e.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
