package main

type IDataStore[T any] interface {
	Load(key T)
	LoadMany()
}
type DataStore[T any] struct {
	IDataStore[T]
	batchLoadFunc func(param T)
}

func NewDataStore[T any](batchLoadFun func(param T)) *DataStore[T] {
	ds := DataStore[T]{}
	ds.batchLoadFunc = batchLoadFun
	return &ds
}

func (ds *DataStore[T]) Load(key T) {
	ds.batchLoadFunc(key)
}
