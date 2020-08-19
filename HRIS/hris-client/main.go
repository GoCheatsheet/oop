package main

import (
	person "oop/person"
)

func main() {
	e := person.New("Sam", "Adolf", 30, 20)
	// e.LeavesRemaining()

	project1 := income.NewFixed("Project 1", 5000)
	project2 := income.NewFixed("Project 2", 10000)
	project3 := income.NewHourly("Project 3", 160, 25)
	incomeStreams := []income.Income{project1, project2, project3}
	income.CalculateNetIncome(e, incomeStreams)
}
