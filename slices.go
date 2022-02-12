package main 

import "fmt"


func main(){
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
	}
	classics := languages[0:3]
	modern := make([]string, 4)
	modern = languages[3:7]
	new := languages[7:9]
	fmt.Printf("classic Languages: %v\n", classics)
	fmt.Printf("Modern Languages: %v\n", modern)
	fmt.Printf("New Languages : %v\n", new)
}
