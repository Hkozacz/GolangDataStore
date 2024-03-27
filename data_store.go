package main

import (
	"sync"
	"time"
)

type IDataStore[T string | int, R any] interface {
	Load(key T, result *chan R, wg *sync.WaitGroup)
	batchLoad()
}

type DataStore[T string | int, R any] struct {
	IDataStore[T, R]
	batchLoadFunc  func(key []T) []R
	isStartedToken chan int
	wg             sync.WaitGroup
	channels       map[T]*chan R
}

func (ds *DataStore[T, R]) batchLoad() {
	defer ds.wg.Done()
	mapLen := len(ds.channels)
	for {
		time.Sleep(100)
		if mapLen == len(ds.channels) {
			ds.isStartedToken = make(chan int, 1)
			ds.isStartedToken <- 1
			break
		} else {
			mapLen = len(ds.channels)
		}
	}

	var results []T
	for key := range ds.channels {
		results = append(results, key)
	}
	batched := ds.batchLoadFunc(results)
	index := 0
	for _, value := range ds.channels {
		*value <- batched[index]
		index++
	}

}

func (ds *DataStore[T, R]) Load(key T, result *chan R, wg *sync.WaitGroup) {
	defer wg.Done()
	_, ok := <-ds.isStartedToken
	if ok {
		close(ds.isStartedToken)
		ds.wg.Add(1)
		go ds.batchLoad()
	}
	ds.channels[key] = result
	ds.wg.Wait()
}

func NewDataStore[T string | int, R any](batchLoadFunc func(key []T) []R) *DataStore[T, R] {
	ds := DataStore[T, R]{}
	ds.batchLoadFunc = batchLoadFunc
	ds.isStartedToken = make(chan int, 1)
	ds.isStartedToken <- 1
	ds.channels = make(map[T]*chan R)
	return &ds
}
