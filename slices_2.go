package main 

import (
	"fmt"
	"reflect"

)

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
        fmt.Printf("New Languages : %v\n\n", new)

	allLangs := languages[:]

	fmt.Println(reflect.TypeOf(allLangs).Kind())

	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]          // length 4 capacity 4
        frameworks = append(frameworks, "Meteor")  // not possible with arrays

        fmt.Printf("all frameworks: %v\n", frameworks)
        fmt.Printf("js frameworks: %v\n", jsFrameworks)
}
