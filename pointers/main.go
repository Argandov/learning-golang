package main

import (
	"fmt"
	"time"
)

type Big struct {
	// a bunch of data to make copying non-trivial
	Data [1_000_000]byte
}

// copyValue takes Big by value (causes a full copy)
func copyValue(b Big) {
	// address of the local copy
	fmt.Printf("  in copyValue, &b = %p\n", &b)
	// mutate the copy
	b.Data[0] = 42
}

// mutatePointer takes *Big and mutates the original
func mutatePointer(b *Big) {
	fmt.Printf("  in mutatePointer, b  = %p\n", b)
	b.Data[0] = 99
}

func main() {
	// create one Big
	orig := Big{}
	fmt.Printf("in main,      &orig = %p\n\n", &orig)

	// 1) Value copy vs. pointer mutation demonstration
	fmt.Println("** Call copyValue(orig) **")
	copyValue(orig)
	fmt.Println("After copyValue, orig.Data[0] =", orig.Data[0], "\n")

	fmt.Println("** Call mutatePointer(&orig) **")
	mutatePointer(&orig)
	fmt.Println("After mutatePointer, orig.Data[0] =", orig.Data[0], "\n")

	// 2) Micro-benchmark: copy vs pointer
	const N = 200
	start := time.Now()
	for i := 0; i < N; i++ {
		copyValue(orig)
	}
	fmt.Printf("copyValue %d× took %v\n", N, time.Since(start))

	start = time.Now()
	for i := 0; i < N; i++ {
		mutatePointer(&orig)
	}
	fmt.Printf("mutatePointer %d× took %v\n", N, time.Since(start))
}
