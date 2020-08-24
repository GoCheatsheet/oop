package main //concurrencycontext

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 3 * time.Second

func contextWithCancel() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		// if n == 6 {
		break
		// }
	}
}

func contextWithDeadline() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second): //Increase the time to invoke ctx.Done()
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}

// func slowOperationWithTimeout(ctx context.Context) (Result, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// 	defer cancel() // releases resources if slowOperation completes before timeout elapses
// 	return slowOperation(ctx)
// }

func contextWithTimeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second): // set this to more than shortDuration
		fmt.Println("Who overslept!!!")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}

func main() {
	contextWithDeadline()
}
