/* RZFeeser | Alta3 Research
   Exploring slices - copy and append       */

package main

import "fmt"
import "time"

func Test() {
	var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond

	startingTime := time.Now().UTC()
	time.Sleep(600 * time.Millisecond)
	endingTime := time.Now().UTC()

	var duration time.Duration = endingTime.Sub(startingTime)

	if duration >= waitFiveHundredMillisections {
		fmt.Printf("Wait %v\nNative [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n",
			waitFiveHundredMillisections,
			duration,
			duration.Nanoseconds()/1e6,
			duration.Seconds())
	}
}

func appendInLoop(elements int) {
	s := make([]int, 0)
	for i := 0; i < elements; i++ {
		s = append(s, i)
	}
}

func writeToIndex(elements int) {
	s := make([]int, elements)
	for i := 0; i < elements; i++ {
		s[i] = i
	}
}

func main() {
	startingTime1 := time.Now().UTC()
	writeToIndex(20000000)
	endingTime1 := time.Now().UTC()
	var duration1 time.Duration = endingTime1.Sub(startingTime1)
	fmt.Printf("Function call duration for writeToIndex in Seconds [%.3f]\n", duration1.Seconds())
	startingTime2 := time.Now().UTC()
	appendInLoop(20000000)
	endingTime2 := time.Now().UTC()
	var duration2 time.Duration = endingTime2.Sub(startingTime2)
	fmt.Printf("Function call duration for appendInLoop in Seconds [%.3f]\n", duration2.Seconds())
}
