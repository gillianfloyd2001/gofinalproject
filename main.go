package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// func createClockIn(writer http.ResponseWriter, request *http.Request) {
// 	body := make([]byte, request.ContentLength)
// 	request.Body.Read(body)

// 	clockIn := Clockin{}           // var name string
// 	json.Unmarshal(body, &clockIn) // body >>>transform>>> clockin

// 	// clockIn := Clockin{EmployeeId: 3, TimeClockedIn: time.Now()}
// 	bytes, _ := json.MarshalIndent(&clockIn, "", "  ")
// 	writer.Write(bytes)
// }

// func createClockOut(writer http.ResponseWriter, request *http.Request) {
// 	body := make([]byte, request.ContentLength)
// 	request.Body.Read(body)

// 	clockOut := Clockout{}
// 	json.Unmarshal(body, &clockOut)

// 	bytes, _ := json.MarshalIndent(&clockOut, "", "  ")
// 	writer.Write(bytes)
// }

func listEmployees(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, Employees())
}

func singleEmployee(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, GetEmployee(id))
}

func createEmployee(c echo.Context) (err error) {
	employee := Employee{}
	if err = c.Bind(employee); err != nil {
		return
	}
	employee.Create()
	return c.JSON(http.StatusCreated, employee)
}

func patchSingleEmployee(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, UpdateEmployee(id))
}

func deleteSingleEmployee(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, DeleteEmployee(id))
}

// func createListEmployees(writer http.ResponseWriter, request *http.Request) {
// 	if request.Method == "GET" {
// 		employees := Employees()

// 		bytes, _ := json.MarshalIndent(&employees, "", "  ")
// 		writer.Write(bytes)
// 	} else if request.Method == "POST" {
// 		body := make([]byte, request.ContentLength)
// 		request.Body.Read(body)

// 		employee := Employee{}
// 		json.Unmarshal(body, &employee)

// 		employee.Create()

// 		bytes, _ := json.MarshalIndent(&employee, "", "  ")

// 		writer.Write(bytes)
// 	}
// }

func listDeliveries(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, Deliveries())
}

func singleDelivery(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, GetDelivery(id))
}

func createDelivery(c echo.Context) (err error) {
	delivery := Delivery{}
	if err = c.Bind(delivery); err != nil {
		return
	}
	delivery.Create()
	return c.JSON(http.StatusCreated, delivery)
}

func patchSingleDelivery(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, UpdateDelivery(id))
}

func deleteSingleDelivery(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, DeleteDelivery(id))
}

func main() {
	// http.HandleFunc("/clockin", createClockIn)
	// http.HandleFunc("/employees", createListEmployees)
	// http.HandleFunc("/clockout", createClockOut)

	e := echo.New()
	e.GET("/employees", listEmployees)
	e.GET("/employee:id", singleEmployee)
	e.POST("/employees", createEmployee)
	e.PATCH("/employee:id", patchSingleEmployee)
	e.DELETE("/employee:id", deleteSingleEmployee)

	e.GET("/deliveries", listDeliveries)
	e.GET("/deliveries:id", singleDelivery)
	e.POST("/deliveries", createDelivery)
	e.PATCH("/deliveries:id", patchSingleDelivery)
	e.DELETE("/deliveries:id", deleteSingleDelivery)

	fmt.Println("Listening on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
