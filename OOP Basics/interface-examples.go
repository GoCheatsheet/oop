pakcage interfaceexample

func mainTypeCheck() {
	// var i interface{}
	// iTest := 10.0
	// i = iTest

	// switch v := i.(type) {
	// case int:
	// 	fmt.Printf("Twice %v is %v\n", v, v*2)
	// case string:
	// 	fmt.Printf("%q is %v bytes long\n", v, len(v))
	// default:
	// 	fmt.Printf("I don't know about type %T!\n", v)
	// }

	var val interface{} = bohoo{}
	myfun(val)

}

type iCheckMe interface {
	sayhey()
}

type bohoo struct{}

func (b bohoo) sayheya() {
	Println("Bohoo")
}

func myfun(a interface{}) {

	// Extracting the value of a
	_, ok := a.(iCheckMe)
	if !ok {
		fmt.Println("Paniced")
	} else {
		// fmt.Println("Value: ", val)
	}

}
