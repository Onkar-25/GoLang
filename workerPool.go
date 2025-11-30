package main

import (
	"fmt"
	"sync"
)

func worker(jobs, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("job received", job)
		result <- job * job
	}
}

func workerPoolunbufferedCh() {
	jobs := make(chan int)
	result := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobs, result, &wg)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func(){
		wg.Wait()
		close(result)
	}()

	for res:= range result{
		fmt.Println("result",res)
	}

}
