package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
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

func singleDelivery(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, GetDelivery(id int))
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

	// http.HandleFunc("/deliveries", createListDeliveries)

	e := echo.New()
	e.GET("/deliveries", listDeliveries)
	e.GET("/deliveries:id", singleDelivery)
	e.POST("/deliveries", createDelivery)
	
	fmt.Println("Listening on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
