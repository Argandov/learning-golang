package main

/*
	0. Declarations vs. Statements
		Statements can't live outside functions, including Short-variable-declarations
	1. Structs,
	2. Modules
*/

import (
	"fmt"
	"modules-learning/cloud"
	"modules-learning/auth"
)

// Declaration:
type User struct{
	name 		string
	id 			int
	active 	bool
	}

func main() {
	cloud.Aws()
	
	user0 := User{"Josh", 12345, true}
	auth.Auth(user0.name)
	
	// Short-variable-declaration == STATEMENT
	user1 := struct{
		name 		string
		id 			int
		active 	bool
	}{
		"Jenny",
		12345,
		true,
	}

	fmt.Println("")
	auth.Auth(user1.name)
}
