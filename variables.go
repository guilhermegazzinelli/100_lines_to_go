package main 

import "fmt"

/* single variable */
var a int

/* Declare a group of variables */
var (
	b bool
	c float32
	d string
)

func main(){
	a = 42
	b, c = true, 32.0
	d = "Aqui vai uma bela String"
	fmt.Println(a, b, c, d)
}

