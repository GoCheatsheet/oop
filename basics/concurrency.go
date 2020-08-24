package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

func sayC(ch chan int) {
	for i := <-ch; i < 6; i++ {
		ch <- i

		// if we do not clean the buffer here,
		// the output will always be 0
		<-ch
	}
	// time.Sleep(100 * time.Millisecond)
}

// Checking how the goroutine behaves when
// multiple values are passed to one channel
func oneBufferAnomaly() {
	c := make(chan int)

	go sayC(c)
	fmt.Println(<-c)

	go sayC(c)
	fmt.Println(<-c)

	go sayC(c)
	fmt.Println(<-c)

	go sayC(c)
	fmt.Println(<-c)

	go sayC(c)
	fmt.Println(<-c)

	go sayC(c)
	fmt.Println(<-c)

}
func sayCap(ch chan int) {
	// cInc := <- ch + 1
	// cc := <- ch
	// fmt.Println( cc )

	ch <- 1
	ch <- 2

	// Extracting the value in goroutine
	// does not create a dead lock
	<-ch
	<-ch
}

func twoBuffer() {
	ch := make(chan int, 2) // Default channel which can hold 1 message.

	go func() {
		ch <- 1
		ch <- 2
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	go func(ch chan int) {
		ch <- 7
		ch <- 8
	}(ch) // Send ch to the function.
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch <- 9 // Error. Cannot pass value to channel no in a goroutine.
	ch <- 10
	v := <-ch // Remove/extract  the value from the from channel ch
	fmt.Println(v)
	fmt.Println(<-ch)

	go sayCap(ch)
}

// chan close example
func fibonacciChanClose(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibClient() {
	c := make(chan int, 10)
	go fibonacciChanClose(cap(c), c)

	// Check if the channel is closed or not; If close throws false
	if out, ok := <-c; ok {
		fmt.Println(out)
	}

	// Will not stop reading data from channel until the channel is closed
	for i := range c {
		fmt.Println(i)
	}

	if out, ok := <-c; ok {
		fmt.Println(out)
	}

	// c <- 21 // Will throw error as the channel is closed
}

// Multiple chan example
func fibonacciNChan(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacciNChanClient() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciNChan(c, quit)
}

// Test with multi chan. Client below.
// The output is random.
func multiChanClientWithBuffer(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// fmt.Printf("select %v \n", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func multiChanClientWithBufferClinet() {
	c := make(chan int)
	quit := make(chan int)

	go func(c chan int) {
		for i := 0; i < 5; i++ {
			<-c
		}
		quit <- 0
		close(c)
	}(c)
	go multiChanClientWithBuffer(c, quit)

	for i := range c {
		if _, ok := <-c; ok {
			fmt.Printf("range %v\n", i)
		}
	}
}

func timeOut() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	c <- 1
	<-o
}

func main() {
	runtime.GOMAXPROCS(4)

	// oneBufferAnomaly()
	// twoBuffer()
	// fibonacciNChanClient()
	// multiChanClientWithBufferClinet()

	timeOut()
}
