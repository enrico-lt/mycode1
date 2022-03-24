/* Alta3 Reseach | RZFeeser
   Goroutines */

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex10  sync.Mutex
	balance1 int
)

func init() {
	balance1 = 1000
}

func deposit1(value int, wg *sync.WaitGroup) {
	//wg.Add(1)
	go depositImpl(value)
	wg.Done()
}

func depositImpl(value int) {
	fmt.Printf("Depositing %d to account with balance1: %d\n", value, balance1)
	mutex10.Lock()
	balance1 += value
	mutex10.Unlock()
}

func withdraw1(value int, wg *sync.WaitGroup) {
	//wg.Add(1)
	go withdrawImpl(value)
	//time.Sleep(time.Second * 3)
	wg.Done()
}

func withdrawImpl(value int) {
	fmt.Printf("Withdrawing %d from account with balance1: %d\n", value, balance1)
	mutex10.Lock()
	balance1 -= value
	mutex10.Unlock()
}

func main() {
	var wg1 sync.WaitGroup
	wg1.Add(3)
	withdraw1(750, &wg1) // if you pass a WaitGroup to a function, use a pointer
	deposit1(1000, &wg1)
	withdraw1(500, &wg1)

	time.Sleep(time.Second)
	wg1.Wait() // wait until we finish

	fmt.Printf("New Balance %d\n", balance1)

	// This doesnt work because the methods are called so fast that the "go" never is invoked and
	// main is finished
	// Wait just wants to see 3 items in wait group, which exist at wg1.Wait()
	// Only works with sleeps
}
