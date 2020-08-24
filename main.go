package main

import (
	"fmt"
	. "fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
)

func frogJumpClient() {
	var want, got int

	// case 1
	want = 3
	got = frogJump(10, 85, 30)
	// 10 + 30 * steps >= 85
	// steps = ( 85 - 10 ) / 30 = 75/30 = 2. Here remainder is more than 1, so increase
	fmt.Printf("1 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("1 Ok")
	}

	// case 2
	want = 8
	got = frogJump(5, 85, 10) // 5 + 10 * 8
	fmt.Printf("2 g:%d	w:%d\n", got, want)
	// 5 + 10 * steps >= 85
	// steps = ( 85 - 5 ) / 10 = 80/10 = 8
	if got == want {
		fmt.Println("2 Ok")
	}

	// case 3 X==Y
	want = 0
	got = frogJump(85, 85, 30)
	fmt.Printf("3 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("3 Ok")
	}

	// case 4
	// Starting point is 0
	want = 10
	got = frogJump(0, 10, 1)
	fmt.Printf("4 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("4 Ok")
	}

	// case 5
	// Jump == Step
	want = 1
	got = frogJump(1, 10, 10)
	fmt.Printf("5 g:%d	w:%d\n", got, want)
	if got == want {
		fmt.Println("5 Ok")
	}

}

// Write a function:
// func Solution(X int, Y int, D int) int
// that, given three integers X, Y and D,
// returns the minimal number of jumps D from
// position X to a position equal to or greater
// than Y.
// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp
func frogJump(X int, Y int, D int) int {

	if X == Y {
		return 0
	}

	if Y == X+D {
		return 1
	}

	mod := (Y - X) % D
	remainder := (Y - X) / D

	if mod > 0 {
		remainder++
	}

	return remainder
}

// Condition:
// Each element of the array can be paired with
// another element that has the same value, except
// for one element that is left unpaired.
// Write a function given an array A consisting of N integers
// returns value of the unpaired elements.
// func Solution(A []int) int

// This is straight forward so will not be solving now
func oddOccurances(A []int) int {
	// Sort A
	// Count values: total odd element t = 0
	// loop range A:
	// element count e = 1
	// if current value is same as previous: increment e
	// else:
	//   check if e is odd: if true then increment t, set e to 1
	// return t

	return 0
}

// Solution scores 100% in Codility Arrays CyclicRotation
// Time O(n), space O(1)
// Resources:
// https://leetcode.com/problems/rotate-array/discuss/806176/20ms-in-C-with-Time-O(n)-space-O(1)-accepted-solution
// https://www.geeksforgeeks.org/reversal-algorithm-right-rotation-array
// https://www.techiedelight.com/right-rotate-an-array-k-times
// https://solution.programmingoneonone.com/2020/07/hackerrank-arrays-left-rotation-solution.html
// https://www.javatpoint.com/java-program-to-right-rotate-the-elements-of-an-array
// https://www.geeksforgeeks.org/python-program-for-program-for-array-rotation-2
// https://www.geeksforgeeks.org/reversal-algorithm-right-rotation-array
// https://stackoverflow.com/questions/21322173/convert-rune-to-int
func cyclicRotation(nums []int, k int) []int {
	numsSize := len(nums)

	if k == 0 || k == numsSize || k > 100 || numsSize == 1 || numsSize == 0 || numsSize > 100 {
		return nums
	}

	if k > numsSize {

		mod := k % numsSize
		fmt.Println(mod)

		if mod == 0 {
			k = k / numsSize
		} else {
			k = mod
		}
	}

	fmt.Println(k)

	rev(nums, 0, numsSize-1)
	rev(nums, 0, k-1)
	rev(nums, k, numsSize-1)

	return nums
}

func rev(a []int, b, e int) {
	for b < e {
		t := a[b]
		a[b] = a[e]
		a[e] = t
		b++
		e--
	}
}

// rotate array as given k times
func cyclicRotationClient() {
	// A := []int{3, 8, 9}
	// K := 2

	// OUT BOUND EXCEPTION
	// Fixing k value fixed it
	// A := []int{3, 8, 9}
	// K := 7

	// OUT BOUND EXCEPTION
	A := []int{3, 8}
	K := 5

	cyclicRotation(A, K)
	fmt.Println(A)

	// brute force ; need more math on this
	// swap with next element swap(k, k+1)
	// if last element, swap with the first element

	// if the rotation == lenght of array, no need to rotate
	// if len == 1 return array

}

// min 1
// max 483647
// Find longest sequence of zeros in binary representation of an integer.
// https://app.codility.com/programmers/lessons/1-iterations
func longestBinaryGap(N int) int {

	if N < 5 {
		return 0
	}

	// x := int(N)
	s := fmt.Sprintf("%b", N)
	fmt.Println(s)
	// c := []int(s)

	m, c := 0, 0
	for _, sv := range s {
		// fmt.Println(sv)

		// max counter m, internal counter c = 0
		// start internal counter ic when sv 49
		// if sv == 1; if c>m m = c; c=0
		// if sv == 0; c++

		if sv == 49 {
			if c > m {
				m = c
			}
			c = 0
		} else {
			c++
		}
		fmt.Printf("D:m:c %d:%d:%d \n", sv, m, c)
	}

	return m
}

func longestBinaryGapClient() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")
	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	fmt.Println(s)

	// 561892
	longestBinaryGap(561892)
}

func Solution(A []int) int {
	// write your code in Go 1.4

	sort.Ints(A)
	minVal := 1

	if len(A) == 1 && (A[0] >= 1000000 || A[0] < 0) {
		return 1
	}

	if A[len(A)-1] <= 0 {
		return 1
	}

	lenA := len(A) - 2
	i := 0
	for i <= lenA {
		if (A[i] >= 0) && (A[i+1]-A[i] > 1) {
			return A[i] + 1
		}
		i++
	}

	if (i-1 == lenA) && (A[i] > minVal) {
		return A[i] + 1
	}

	return minVal
}
func SolutionCllinet() {
	caseOne()
	caseTwo()
	caseThree()
	caseFour()
	caseFive()
}
func caseFive() {
	a := []int{-1, -2000, -3}
	ret := Solution(a)

	if ret == 1 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseFour() {
	a := []int{-1, 000, 000}
	ret := Solution(a)

	if ret == 1 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseThree() {

	a := []int{1, 3, 6, 4, 1, 2, 10000000}
	ret := Solution(a)

	if ret == 5 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseTwo() {
	a := []int{9, 3, 6, 1, 15, 21, 8, 36, 45, 5}
	ret := Solution(a)

	if ret == 2 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}
func caseOne() {
	a := []int{-1, -50, 1, 2, 3}
	ret := Solution(a)

	if ret == 4 {
		fmt.Println("ok")
	} else {
		fmt.Printf("ISSUE %v", ret)
	}
}

func sliceExample() {
	ss := "a b c Sas Sa"
	fmt.Println(strings.Count(ss, "Sa"))

	// search()
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	for _, str := range strSlice {
		if str == "Italya" {
			fmt.Println(str)
		}
	}
	// fmt.Println(itemExists(strSlice, "Canada"))
	// fmt.Println(itemExists(strSlice, "Africa"))
}

func testRef(s string) {
	s = "c" + s[1:] // you cannot change string values by index, but you can get values instead.
}

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

// ByName implements sort.Interface based on the Age field.
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func sortExample() {
	family := []Person{
		{"ZAlice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByName(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}

func search() {
	// srcStr := "aaa"
	// dstStr := ""
	// fmt.Println(copy(dstStr, srcStr)

	a := []int{9, 3, 6, 1, 15, 21, 8, 36, 45, 5}
	refa := a       // another way to copy a[:len(a)] // copy by reference
	sort.Ints(refa) // so this also sorts a but we want to keep the original indexing.

	// remove(refa, 8)
	// var b []int
	// capa := len(a)
	// b = a[:capa]

	// var arr [7]int
	// arr = a
	// fmt.Println("arr original %v\n", arr)

	b := make([]int, len(a), cap(a))
	copy(b, a)

	// search string in slice

	x := 6
	i := sort.Search(len(b), func(i int) bool { return b[i] >= x })
	if i < len(a) && b[i] == x {
		fmt.Printf("original %v\n", a)
		fmt.Printf("found %d at index %d in %v\n", x, i, b)
	} else {
		fmt.Printf("%d not found in %v\n", x, b)
	}
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func xmain() {

	family := []struct {
		Name string
		Age  int
	}{
		{"ZAlice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// sort.SliceStable(family)
	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)

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
