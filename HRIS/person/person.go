package person

import (
	"fmt"
)

type Person struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

type Employee interface {
	int CalculateNetIncome()
}

// NewPerson Instantiate new person
func NewPerson(firstName string, lastName string, totalLeave int, leavesTaken int) Person {
	return Person{firstName, lastName, totalLeave, leavesTaken}
}

func (e Person) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
