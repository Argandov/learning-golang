package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	fmt.Printf("---- Start Program at %v\n", time.Now())
	fmt.Println("| ARCH:", runtime.GOARCH)
	fmt.Println("| OS:", runtime.GOOS)
	fmt.Println("| CPUs:", runtime.NumCPU())
	fmt.Println("| Goroutines:", runtime.NumGoroutine())
	fmt.Println("----------------------------")
}

func worker(id string) {
	fmt.Println("> worker", id, "started")
	time.Sleep(time.Second)
	fmt.Println("_ worker", id, "done")
}

func main() {
	start := time.Now()

	id := "_goroutine"
	go worker(id)
	id = "no_goroutine"
	worker(id)

	end := time.Now()
	fmt.Println("---- End Program. Duration:", end.Sub(start))
	fmt.Println("| Goroutines:", runtime.NumGoroutine())
}
