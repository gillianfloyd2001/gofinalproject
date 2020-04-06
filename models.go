package main

import (
	"time"
)

// This creates an Employee
type Employee struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

// This creates a ClockIn
type Clockin struct {
	ID            uint      `json: "id"`
	EmployeeID    *Employee `json:"employeeId"`
	TimeClockedIn time.Time `json:"timeClockedIn"`
}

// This creates a clockout
type Clockout struct {
	ID             uint      `json: "id"`
	EmployeeId     *Employee `json: "employeeId"`
	Timeclockedout time.Time `json: "timeClockedOut"`
}

// This creates a Delivery
type Delivery struct {
	ID uint `json: "id"`
}

// This creates a Payout
type Payout struct {
	ID         uint `json: "id"`
	EmployeeID *Employee
	Timestamp  time.Time
	Amount     uint
}
