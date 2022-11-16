package main

import (
	"fmt"
	"sync"
)

const N = 50

func main() {

	in1, in2, out := make(chan int, 1), make(chan int, 1), make(chan int, N)

	merge2Channels(fn, in1, in2, out, N)

	for i := 0; i < N; i++ {
		in1 <- i
		in2 <- i

	}

	for i := 0; i < N; i++ {

		fmt.Println(<-out)

	}

}

func fn(val int) int {

	return val * 2

}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		var wg sync.WaitGroup
		results := make([]int, n)
		wg.Add(n * 2)
		for i := 0; i < n; i++ {
			val1 := <-in1
			val2 := <-in2
			go func(num int, val1 int) {
				defer wg.Done()
				results[num] = results[num] + fn(val1)
			}(i, val1)
			go func(num int, val2 int) {
				defer wg.Done()
				results[num] = results[num] + fn(val2)
			}(i, val2)
		}
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- results[i]
		}
	}()

}
