package polymorphismexample

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Override default Print
func (h Human) String() string {
	return "Name:" + h.name + ", Age:" + strconv.Itoa(h.age) + " years, Contact:" + h.phone
}

// method
func (h Human) SayHi() {
	h.name = "Anit"
	Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// method
func (h Human) Sing(lyrics string) {
	Println("La la la la...", lyrics)
}

// method
func (e Employee) SayHi() {
	Printf("Employee Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func (e Employee) Sing(lyrics string) {
	Println("Employee sings La la la la...", lyrics)
}

func client() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}
	rina := Employee{Human{"Rina", 22, "444-222-XXX"}, "Google.", 11000}

	// define interface i
	var i Men

	//i can store Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i can store Employee
	// i = tom
	// fmt.Println("This is Tom, an Employee:")
	// i.SayHi()
	// i.Sing("Born to be wild")

	go rina.SayHi()

	// slice of Men
	Println("Let's use a slice of Men and see what happens")
	mensList := make([]Employee, 2)
	// these three elements are different types but they all implemented interface Men
	mensList[0], mensList[1] = tom, sam

	for _, men := range mensList {
		Println("This Human is : ", men)
		go men.SayHi()
	}
}
