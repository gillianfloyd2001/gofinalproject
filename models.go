package main

import (
	"time"
)

// This creates an Employee
type Employee struct {
	Id       uint   `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Position string `json:"position" form:"position" query:"position"`
}

// This creates a ClockIn
type Clockin struct {
	Id            uint      `json:"id" form:"id" query:"id"`
	EmployeeId    *Employee `json:"employeeId" form:"employeeId" query:"employeeId"`
	TimeClockedIn time.Time `json:"timeClockedIn" form:"timeClockedIn" query:"timeClockedIn"`
}

// This creates a clockout
type Clockout struct {
	Id             uint      `json:"id" form:"id" query:"id"`
	EmployeeId     *Employee `json:"employeeId"  form:"employeeId" query:"employeeId"`
	Timeclockedout time.Time `json:"timeClockedOut" form:"timeClockedOut" query:"timeClockedOut"`
}

// This creates a Delivery
type Delivery struct {
	Id          uint   `json:"id"`
	Name        string `json: "name"`
	PhoneNumber string `json: "phoneNumber"`
	Address     string `json:"address"`
	Tip         bool   `json:"tip"`
}

// This creates a Payout
type Payout struct {
	Id         uint      `json: "id"`
	EmployeeId *Employee `json: "employeeid"`
	Timestamp  time.Time `json: "timeStamp"`
	Amount     uint      `json: "amount"`
}
