package main

import (
	"fmt"
	"sync"
)

func batchLoadFunc(keys []string) []int {
	var res []int
	fmt.Println("once?")
	for _, element := range keys {
		res = append(res, len(element))
	}
	return res
}

func main() {
	ds := NewDataStore[string, int](batchLoadFunc)
	var wg sync.WaitGroup
	result := make(chan int, 100)
	result2 := make(chan int, 100)
	wg.Add(1)
	go ds.Load("key", &result, &wg)
	wg.Add(1)
	go ds.Load("key2", &result2, &wg)
	wg.Wait()
	fmt.Println(<-result)
	fmt.Println(<-result2)
}
