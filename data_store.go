package main

type IDataStore interface {
	Load()
	LoadMany()
}
type DataStore struct {
	IDataStore
	batchLoadFunc func(param string)
}

func NewDataStore(batchLoadFun func(param string)) *DataStore {
	ds := DataStore{}
	ds.batchLoadFunc = batchLoadFun
	return &ds
}

func (ds *DataStore) Load() {
	ds.batchLoadFunc("Hello World")
}
