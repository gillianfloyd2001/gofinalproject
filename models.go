package main

import (
	"time"
)

// This creates an Employee
type Employee struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

// This creates a ClockIn
type Clockin struct {
	Id            uint      `json:"id"`
	EmployeeId    *Employee `json:"employeeId"`
	TimeClockedIn time.Time `json:"timeClockedIn"`
}

// This creates a clockout
type Clockout struct {
	Id             uint      `json:"id"`
	EmployeeId     *Employee `json:"employeeId"`
	Timeclockedout time.Time `json:"timeClockedOut"`
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
	Id         uint `json: "id"`
	EmployeeId *Employee
	Timestamp  time.Time
	Amount     uint
}
