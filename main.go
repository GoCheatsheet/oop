package main

import (
	. "fmt"
)

func testRef(s string) {
	s = "c" + s[1:] // you cannot change string values by index, but you can get values instead.
}

func main() {
	// mainBrokenInterface()
	// var a int8 = 11

	// var b int32

	// b = int32(a)

	// vname1, vname2, vname3 := "v1", "v2", 1
	// Printf("Let's use a slice of Men and see what happens %s %s %d %d \n", vname1, vname2, vname3, b)

	// c := make(chan int, 1) // change 2 to 1 will have runtime error, but 3 is fine
	// c <- 1
	// c <- 2
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// s := "hello this "
	// c := []byte(s) // convert string to []byte type
	// c[0] = 'c'
	// s2 := string(c) // convert back to string type
	// fmt.Printf("%s\n", s2)

	// s := "hello"
	// s = "c" + s[1:] // you cannot change string values by index, but you can get values instead.
	// Printf("%s\n", s)

	// err := errors.New("emit macho dwarf: elf header corrupted")
	// if err != nil {
	// 	Println(err)
	// }

	// bytes := [5]byte{'h', 'e', 'h', 'e', 'h'}
	// Printf("%s\n", bytes)

	// const (
	// 	x = iota // x == 0
	// 	y        //= iota // y == 1
	// 	z        //= iota // z == 2
	// 	w        // If there is no expression after the constants name, it uses the last expression,
	// 	//so it's saying w = iota implicitly. Therefore w == 3, and y and z both can omit "= iota" as well.
	// )

	// Printf("%d %d %d %d ", x, y, z, w)

	// s := "hello"
	// s = "c" + s[1:] // you cannot change string values by index, but you can get values instead.
	// fmt.Printf("%s\n", s)

	// c := [...]int{4, 5, 6} // use `…` to replace the length parameter and Go will calculate it for you.

	// arr := []byte{'a', 'b', 'c', 'd', 'e'}
	// for k, v := range arr {
	// 	Printf("%v:%v\n", k, string(v))
	// }

	// testRef(string(arr))
	// Printf("%v \n", string(arr))

	// aslice := arr[:]
	// aslice = append(aslice, 'k')
	// fmt.Printf("%d\n", cap(aslice))

	// rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// csharpRating, keyExists := rating["C#"] // ok is false as C# does not exists,

	// Printf("%f %s \n", csharpRating, keyExists)
	// strOrg := "Check"
	// isTrue := true
	// Printf("%v %v \n", strOrg, isTrue)

	var mapList = make(map[int]int)
	mapList[1] = 1
	mapList[2] = 2

	// var mapEg = map[int]int{1: 1, 2: 2}
	for k, v := range mapList {
		Printf("%v:%v\n", k, v)
	}

	mapList = nil

	Printf("%T\n", mapList)
}
