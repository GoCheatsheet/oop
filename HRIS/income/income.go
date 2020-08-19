package income

import (
	"fmt"
	"oop/person"
)

type IncomeSource interface {
	calculate() int
	source() string
}

type Income struct {
	person.Person
	netincome int
}

// NewIncome instantiates Income
// Associate income with employee
func NewIncome(employeeObj person.Person) *Income {
	return &Income{
		Employee: employeeObj,
	}
}

// CalculateNetIncome calculates the net income of an employee
func (income Income) CalculateNetIncome(employee *employee.Employee, ic []IncomeSource) {
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		income.netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}
