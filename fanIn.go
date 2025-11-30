package main

import (
	"fmt"
	"sync"
)

func generator(start int, ch chan int) {
	for i := start; i < start+3; i++ {
		ch <- i
	}
	close(ch)
}

func fanIn(ch1, ch2 chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(2)
	go func() {

		defer wg.Done()

		for c := range ch1 {
			out <- c
		}
	}()

	go func() {
		defer wg.Done()
		for c := range ch2 {
			out <- c
		}
	}()

	return out
}

func fanInGen() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go generator(10, ch1)
	go generator(100, ch2)
	var wg sync.WaitGroup
	merged := fanIn(ch1, ch2, &wg)

	go func() {
		wg.Wait()
		close(merged)
	}()

	for merge := range merged {
		fmt.Println(merge)
	}
	//for i := 0; i < 6; i++ {
	//fmt.Println(<-merged)
	//}
}
