package cloud

import "fmt"

func Aws() {
	fmt.Println("aws")
}

// a non-exported function

func aws1() {
	fmt.Println("aws")
}
