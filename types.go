package main 

import "fmt"

func main(){
	// user specified types

	const a int32 = 12
	const b float32 = 20.5
	var c complex128 = 1 + 4i
	var d uint16 = 14


	n := 42
	pi := 3.14
	x, y := true, false
	z := "Go is SUPER"

	fmt.Printf("User-Specified types:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("Default Types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}
