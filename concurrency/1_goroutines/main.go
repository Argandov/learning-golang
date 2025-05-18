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

func big_timer() {
	big_time := 5
	for i := 0; i <= big_time; i++ {
		fmt.Println("BigTimer; i=", i+0)
		time.Sleep(time.Second)
	}
}

func small_timer() {
	small_time := 2
	for i := 0; i <= small_time; i++ {
		fmt.Println("SmallTimer; i=", i+1)
		time.Sleep(time.Second)
	}
}

func worker(id string) {
	for i := 0; i < 5; i++ {
		fmt.Println(id, ":", i+1, "/5")
		time.Sleep(time.Second)
	}
}


func main() {
	/* 

		main() == goroutine G0
		big_timer() == goroutine G1
		go worker() == goroutine G2

	G1/G2 runs dependent on G0. If G0 stops, G1/G2 halt.

	*/
	start := time.Now()
	
	go big_timer()
	small_timer()

	// G = Goroutine
	go worker("[G] Worker 1")
	worker ("_ Worker N 2")

	end := time.Now()
	fmt.Println("---- End Program. Duration:", end.Sub(start))
	fmt.Println("| Goroutines:", runtime.NumGoroutine())
}
