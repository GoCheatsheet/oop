package main

import (
    "fmt"
    "runtime"
    // "time"
)

func sayC(ch chan int) {
    for i := 0; i < 6; i++ {
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

func main() {
    runtime.GOMAXPROCS(4)

    oneBufferAnomaly()
    

}