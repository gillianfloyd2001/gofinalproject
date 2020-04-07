package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createClockIn(writer http.ResponseWriter, request *http.Request) {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)

	clockIn := Clockin{}           // var name string
	json.Unmarshal(body, &clockIn) // body >>>transform>>> clockin

	// clockIn := Clockin{EmployeeId: 3, TimeClockedIn: time.Now()}
	bytes, _ := json.MarshalIndent(&clockIn, "", "  ")
	writer.Write(bytes)
}

func createClockOut(writer http.ResponseWriter, request *http.Request) {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)

	clockOut := Clockout{}
	json.Unmarshal(body, &clockOut)

	bytes, _ := json.MarshalIndent(&clockOut, "", "  ")
	writer.Write(bytes)
}

func createListEmployees(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		employees := Employees()

		bytes, _ := json.MarshalIndent(&employees, "", "  ")
		writer.Write(bytes)
	} else if request.Method == "POST" {
		body := make([]byte, request.ContentLength)
		request.Body.Read(body)

		employee := Employee{}
		json.Unmarshal(body, &employee)

		employee.Create()

		bytes, _ := json.MarshalIndent(&employee, "", "  ")

		writer.Write(bytes)
	}
}

func createListDeliveries(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		deliveries := Deliveries()

		bytes, _ := json.MarshalIndent(&deliveries, "", "  ")
		writer.Write(bytes)
	} else if request.Method == "POST" {
		body := make([]byte, request.ContentLength)
		request.Body.Read(body)

		delivery := Delivery{}
		json.Unmarshal(body, &delivery)

		delivery.Create()

		bytes, _ := json.MarshalIndent(&delivery, "", "  ")

		writer.Write(bytes)
	}
}

func main() {
	http.HandleFunc("/clockin", createClockIn)
	http.HandleFunc("/employees", createListEmployees)
	http.HandleFunc("/deliveries", createListDeliveries)
	http.HandleFunc("/clockout", createClockOut)

	fmt.Println("Listening on 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
