package main

import (
	"fmt"
	"runtime"
	"time"

	uuid "github.com/gofrs/uuid"
)

const (
	numElements = 10000000
)

var foo = map[string]int{}

func timeGC() {
	t := time.Now()
	runtime.GC()
	fmt.Printf("gc took: %s\n", time.Since(t))
}

func main() {
	for i := 0; i < numElements; i++ {
		u := uuid.Must(uuid.NewV4()).String()
		foo[u] = i
	}

	for {
		timeGC()
		time.Sleep(1 * time.Second)
	}
}
