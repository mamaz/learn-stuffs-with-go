package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Composition: Behaviour
// ============================
type Person struct {
	ID string
	Leg
	Hand
}

type Leg struct {
}

func (l *Leg) Walk() {
	fmt.Println("walk")
}

type Hand struct {
}

func (l *Leg) Punch() {
	fmt.Println("wave")
}

// Composition: fields
// ============================

type Employee struct {
	ID     string
	Name   string
	Salary int
}

type Manager struct {
	Employee
	Staffs    []Employee
	Allowance int
}

type VP struct {
	Employee
	Staffs []Manager
	Stocks int
}

func main() {
	p := Person{ID: "12"}
	p.Leg = Leg{}
	p.Hand = Hand{}

	// person can walk, and can punch
	p.Walk()
	p.Punch()

	// example: Employee hierarchy
	staff := Employee{
		ID:     "32",
		Name:   "Carlos de Melo",
		Salary: 1000000,
	}
	fmt.Printf("staff: %+v\n", staff)
	fmt.Println("")

	m := Manager{
		Employee: Employee{
			ID:     "1",
			Name:   "Hisma",
			Salary: 10000000,
		},
		Staffs: []Employee{
			{
				ID:     "2",
				Name:   "Redouan Barkawi",
				Salary: 5000000,
			},
			{
				ID:     "3",
				Name:   "Bejo Sugiantoro",
				Salary: 5000000,
			},
		},
		Allowance: 1000,
	}

	fmt.Printf("Manager: %+v\n", m)

	managerJSON, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Manager JSON: %+v\n", string(managerJSON))
	fmt.Println("")

	vp := VP{
		Employee: Employee{
			ID:     "13",
			Name:   "Ilham",
			Salary: 1000000000,
		},
		Staffs: []Manager{m},
		Stocks: 1000,
	}
	vpJSON, err := json.Marshal(vp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VP: %+v\n", vp)
	fmt.Printf("VP JSON: %+v\n", string(vpJSON))
}
