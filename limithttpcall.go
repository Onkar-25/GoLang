package main

import (
	"fmt"
	"sync"
	"time"
)

func callApi(id int, limit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	limit <- struct{}{}
	fmt.Println("start id",id)
	time.Sleep(500*time.Millisecond)
	fmt.Println("done id :",id)
	<- limit
}

func limitHttpCall(){
	limit := make(chan struct{},3)
	var wg sync.WaitGroup
	for i:=0;i<=10;i++{
		wg.Add(1)
		go callApi(i,limit,&wg)
	}
	wg.Wait()
}