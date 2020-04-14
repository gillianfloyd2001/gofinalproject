package main

import (
	"fmt"
	"net/http"

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

func listDeliveries(c echo.Context) {
	return c.JSON(http.StatusOK, Deliveries())
}

func singleDelivery(c echo.Context) {
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

func main() {
	// http.HandleFunc("/clockin", createClockIn)
	// http.HandleFunc("/employees", createListEmployees)
	// http.HandleFunc("/clockout", createClockOut)

	e := echo.New()
	e.GET("/employees", listEmployees)
	e.GET("/employees:id", singleEmployee)
	e.POST("/employees", createEmployee)

	e.GET("/deliveries", listDeliveries)
	e.GET("/deliveries:id", singleDelivery)
	e.POST("/deliveries", createDelivery)

	fmt.Println("Listening on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
