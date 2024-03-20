package main

import "fmt"

func batchLoad(foo string) {
	fmt.Println(foo)
}

func main() {
	ds := NewDataStore(batchLoad)
	ds.Load("Hello world")
}
