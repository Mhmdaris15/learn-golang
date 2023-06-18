package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

func increment() {
	counter = counter + 1
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Counter:", counter)
}

/*

In the above code, we have two goroutines, increment(), that concurrently increment a shared variable counter. 
A race condition occurs when multiple goroutines access and modify a shared resource concurrently without proper synchronization.

In this case, the race condition arises due to the non-atomic nature of the counter increment operation. 
When both goroutines execute counter = counter + 1 simultaneously, they read the current value of counter, perform the increment operation, and write the updated value back. 
However, if the timing is such that one goroutine reads the value before the other goroutine writes its updated value, the final value of counter will be incorrect.

For example, let's say counter starts at 0. Both goroutines read the value as 0, increment it to 1, and attempt to write back the updated value simultaneously. 
However, due to the race condition, one goroutine's write operation may overwrite the update made by the other goroutine, leading to an incorrect final value of counter. 
As a result, the output may vary unpredictably, and the value of counter may not be 2 as expected.

To avoid this race condition, synchronization mechanisms like mutexes or channels can be used to ensure exclusive access to shared resources, 
ensuring that only one goroutine modifies the counter variable at a time.

*/