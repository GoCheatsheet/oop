package main

import (
    "fmt"
    "runtime"
    // "time"
)

func sayC(ch chan int) {
    for i := <- ch; i < 6; i++ {
        ch <- i

        // if we do not clean the buffer here,
        // the output will always be 0
        <- ch 
    }
    // time.Sleep(100 * time.Millisecond)
}

// Checking how the goroutine behaves when
// multiple values are passed to one channel
func oneBufferAnomaly() {
    c := make(chan int)
  
    go sayC(c)
    fmt.Println( <- c )

    go sayC(c)
    fmt.Println( <- c )

    go sayC(c)
    fmt.Println( <- c )

    go sayC(c)
    fmt.Println( <- c )

    go sayC(c)
    fmt.Println( <- c )

    go sayC(c)
    fmt.Println( <- c )

}
func sayCap(ch chan int) {
    // cInc := <- ch + 1
    // cc := <- ch 
    // fmt.Println( cc )

    ch <- 1
    ch <- 2

    // Extracting the value in goroutine
    // does not create a dead lock
    <- ch 
    <- ch
}

func twoBuffer() {
    ch := make(chan int, 2) // Default channel which can hold 1 message.

    go func() { 
        ch <- 1  
        ch <- 2
    }()
    fmt.Println( <- ch )
    fmt.Println( <- ch )

    go func(ch chan int) { 
        ch <- 7  
        ch <- 8
    }(ch)  // Send ch to the function.
    fmt.Println( <- ch )
    fmt.Println( <- ch )

    ch <- 9  // Error. Cannot pass value to channel no in a goroutine.
    ch <- 10
    v := <- ch // Remove/extract  the value from the from channel ch
    fmt.Println( v )
    fmt.Println( <- ch )

    go sayCap(ch)
}

func main() {
    runtime.GOMAXPROCS(4)

    // oneBufferAnomaly()
    twoBuffer()

}