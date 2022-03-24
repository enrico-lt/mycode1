/* Alta3 Reseach | RZFeeser
   Goroutines */

package main

import (
	"fmt"
	"sync"
)

var (
	mutex1   sync.Mutex
	balance1 int
)

func init() {
	balance1 = 1000
}

func deposit1(value int, wg *sync.WaitGroup) {
	wg.Add(1)
	go depositImpl(value)
	wg.Done()
}

func depositImpl(value int) {
	fmt.Printf("Depositing %d to account with balance1: %d\n", value, balance1)
	balance1 += value
}

func withdraw1(value int, wg *sync.WaitGroup) {
	wg.Add(1)
	go withdrawImpl(value)
	wg.Done()
}

func withdrawImpl(value int) {
	fmt.Printf("Withdrawing %d from account with balance1: %d\n", value, balance1)
	balance1 -= value
}

func main() {
	var wg1 sync.WaitGroup

	withdraw1(750, &wg1) // if you pass a WaitGroup to a function, use a pointer
	deposit1(1000, &wg1)
	withdraw1(500, &wg1)
	wg1.Wait() // wait until we finish

	fmt.Printf("New Balance %d\n", balance1)
}
